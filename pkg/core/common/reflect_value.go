package common

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/deemount/gobpmn/pkg/config"
	"github.com/deemount/gobpmn/pkg/core/foundation"
)

// ReflectValue represents the core structure for BPMN reflection
type ReflectValue struct {
	config.ModelConfig
	config.ProcessConfig
	config.ParticipantConfig
	FlowNeighbors    map[string]map[string]interface{}
	InstanceNumField int
}

// newReflectValue ...
func NewReflectValue(model interface{}) (*ReflectValue, error) {
	if model == nil {
		return nil, fmt.Errorf("model cannot be nil")
	}

	// reflect the type of the struct
	typeOf := reflect.TypeOf(model)

	return &ReflectValue{
		ModelConfig: config.ModelConfig{
			Name:     extractPrefixBeforeProcess(typeOf.Name()),
			Fields:   reflect.VisibleFields(typeOf),
			Instance: reflect.New(typeOf).Elem(), // create a new object of the struct
		},
	}, nil

}

// initialize sets up the initial configuration for the ReflectValue
func (v *ReflectValue) initialize() error {
	if err := v.setupInstanceFields(); err != nil {
		return fmt.Errorf("failed to setup target fields: %w", err)
	}

	if err := v.setupDefinition(); err != nil {
		return fmt.Errorf("failed to setup definition: %w", err)
	}

	return nil

}

// setupInstanceFields initializes the target fields
func (v *ReflectValue) setupInstanceFields() error {
	v.InstanceNumField = v.Instance.NumField()

	v.Def = targetFieldName(v, "Def")
	if !v.Def.IsValid() {
		return fmt.Errorf("definition field is not valid")
	}

	return nil

}

// setupDefinition configures the definition field
func (v *ReflectValue) setupDefinition() error {
	if !v.Def.CanSet() {
		return nil
	}

	definitions := foundation.NewDefinitions()
	definitions.SetDefaultAttributes()
	v.Def.Set(reflect.ValueOf(definitions))

	return nil

}

// handleModelType determines and handles the model type (pool or single)
func (v *ReflectValue) handleModelType(q *Quantity, m *Mapping) error {
	if len(m.Anonym) > 0 {
		return v.handlePool(q, m)
	}
	return v.handleSingle(q, m)
}

// handlePool assigns the field Pool to v.Pool.
// It walks through the anonymous fields in Mapping and checks
// if the field contains the word "Pool".
// A pool can only be set once in the BPMN model.
func (v *ReflectValue) handlePool(q *Quantity, m *Mapping) error {
	for _, field := range m.Anonym {
		if strings.Contains(field, "Pool") {
			v.Pool = targetFieldName(v, field)
			q.countFieldsInPool(v)
			if !v.Pool.IsValid() {
				return fmt.Errorf("invalid pool field: %s", field)
			}
			q.countFieldsInInstance(v)
			q.Pool++
			break
		}
	}
	return nil
}

// handleSingle appends the process name to v.ProcessName.
// The process name is extracted in newReflectValue, which is
// the name of the data structure.
// A single process is a process without a pool.
func (v *ReflectValue) handleSingle(q *Quantity, m *Mapping) error {
	for _, bpmnType := range m.BPMNType {
		if strings.Contains(bpmnType, "Process") {
			v.ProcessName = append(v.ProcessName, v.Name)
			q.countFieldsInInstance(v)
			q.Process++
			return nil
		}
	}
	return fmt.Errorf("no process type found in model")
}

// instance holds the data structure of the BPMN model in the ReflectValue.
// the return value is then as a return value in the initial function.
// It is used as a blueprint for the BPMN model, which is then reflected into v.Def.
// It distribute hash values and set the isExecutable field to true.
// It uses the type declaration in the data structure to generate the hash value.
func (v *ReflectValue) instance(m *Mapping) any {
	if len(m.Anonym) > 0 {
		for _, anonymField := range m.Anonym {
			v.anonym(anonymField)
		}
	} else {
		v.nonAnonym(m)
	}

	return v.Instance.Interface()
}

// reflectProcess reflects the processes in the BPMN model.
// v.Process[index of process] is a slice of reflect.Value, which represents the processes in the BPMN model.
// v.Process[0] is the first process in the model and it's used also in a single process arrangement.
func (v *ReflectValue) reflectProcess(q *Quantity) error {
	v.Process = make([]reflect.Value, q.Process)
	method := v.Def.MethodByName("SetProcess")
	if !method.IsValid() {
		return fmt.Errorf("SetProcess method not found")
	}
	results := method.Call([]reflect.Value{reflect.ValueOf(q.Process)})
	if len(results) > 0 && !results[0].IsNil() {
		return fmt.Errorf("SetProcess call failed")
	}
	for processIdx := 0; processIdx < q.Process; processIdx++ {
		v.Process[processIdx] = v.Def.MethodByName("GetProcess").Call([]reflect.Value{reflect.ValueOf(processIdx)})[0]
	}
	return nil
}

// processingWithContext processes the BPMN model with a cancellable context.
func (v *ReflectValue) processingWithContext(ctx context.Context, q *Quantity) error {
	errChan := make(chan error, 1)
	done := make(chan struct{})
	// Run the processing in a goroutine
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			// cleanup, if context is cancelled
			defer close(done)
		default:
			errChan <- v.processing(ctx, q)
		}
	}(context.Background())
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		<-done // Wait for goroutine to finish
		return ctx.Err()
	}
}

// process sets the maps and process parameters in the BPMN model.
// If the pool is greater than 0, then process multiple processes.
// It considers the pool as a collaboration.
// It creates a cancellable context for the processor.
func (v *ReflectValue) processing(ctx context.Context, q *Quantity) error {
	// initialize the slices
	v.ProcessExec = make([]bool, q.Process)
	v.ProcessType = make([]string, q.Process)
	v.ProcessHash = make([]string, q.Process)
	v.ParticipantHash = make([]string, q.Participant)

	// create a new element processor
	processor, err := NewElementProcessor(v, q)
	if err != nil {
		return fmt.Errorf("failed to create element processor: %w", err)
	}

	// create a cancellable context
	processorCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// a pool is detected
	if q.Pool > 0 {

		err := v.configurePool()
		if err != nil {
			return fmt.Errorf("Error configuring pool: %v", err)
		}
		v.multipleProcess(q)
		if err := processor.ProcessElements(processorCtx); err != nil {
			return err
		}

	} else {

		if err := v.configureStandalone(); err != nil {
			return fmt.Errorf("Error processing standalone: %v", err)
		}
		v.standalone(0, q)
		if err := processor.ProcessStandalone(processorCtx); err != nil {
			return err
		}

	}
	return nil
}

// finalizeModel completes the model construction
func (v *ReflectValue) finalizeModel(ctx context.Context, q *Quantity) error {
	if err := v.processingWithContext(ctx, q); err != nil {
		return fmt.Errorf("failed to process: %w", err)
	}

	if err := v.collaboration(q); err != nil {
		return fmt.Errorf("failed to setup collaboration: %w", err)
	}

	return nil

}

// nonAnonym sets the IsExecutable field to true and populates the reflection fields with hash values.
// Note: the method is used in a single process.
func (v *ReflectValue) nonAnonym(m *Mapping) {
	v.setIsExecutableByField(m.Config)
	v.populateReflectionFields(m.BPMNType)
	v.FlowNeighbors = collectFromFieldsWithNeighbors(v.Instance.Interface())
}

// setIsExecutableByField configures the process executable status.
// It sets the IsExecutable field to true.
// Note: maybe redudant; look at method config.
func (v *ReflectValue) setIsExecutableByField(fields map[int]string) {
	for _, field := range fields {
		targetFieldName(v, field).SetBool(true)
		break
	}
}

// setIsExecutableByMethod configures the process executable status.
// It calls SetIsExecutable method on the process.
func (v *ReflectValue) setIsExecutableByMethod(process reflect.Value, isExecutable bool) error {
	method := process.MethodByName("SetIsExecutable")
	if !method.IsValid() {
		return fmt.Errorf("SetIsExecutable method not found")
	}
	method.Call([]reflect.Value{reflect.ValueOf(isExecutable)})
	return nil
}

// populateReflectionFields populates the reflection fields with hash values.
// Note: this method is used in a single process and in usage in the method "nonAnonym".
func (v *ReflectValue) populateReflectionFields(reflectionFields map[int]string) {
	for _, field := range reflectionFields {
		f := targetFieldName(v, field)
		fName, _ := v.Instance.Type().FieldByName(field)
		typ := typ(fName.Name)
		hash, _ := hash(typ)
		f.Set(reflect.ValueOf(hash))
	}
}

// setProcessIdentifiers sets the process ID and name
func (v *ReflectValue) setProcessIdentifiers(process reflect.Value, name string, field reflect.Value) error {
	hash := field.FieldByName("Hash").String()

	// set process ID
	if err := callMethod(process, "SetID",
		[]reflect.Value{
			reflect.ValueOf(strings.ToLower(name)),
			reflect.ValueOf(hash),
		}); err != nil {
		return fmt.Errorf("failed to set process ID: %w", err)
	}

	// set process name
	if err := callMethod(process, "SetName",
		[]reflect.Value{reflect.ValueOf(name)}); err != nil {
		return fmt.Errorf("failed to set process name: %w", err)
	}

	return nil
}

// standalone configures and initializes a single BPMN process.
// It sets up the process properties and applies the required BPMN elements.
// Note: this method is used for a single process and in usage in the method "process".
func (v *ReflectValue) standalone(processIdx int, q *Quantity) error {
	if err := v.configureStandalone(); err != nil {
		return fmt.Errorf("failed to configure process: %w", err)
	}

	if err := v.applyProcessMethods(processIdx, q); err != nil {
		return fmt.Errorf("failed to apply process methods: %w", err)
	}

	return nil

}

// configureStandalone configures a single BPMN process.
func (v *ReflectValue) configureStandalone() error {
	if v.Process == nil || len(v.Process) == 0 {
		return fmt.Errorf("invalid process: Process slice is nil or empty")
	}

	process := v.Process[0]

	for i := 0; i < v.InstanceNumField; i++ {
		field := v.Instance.Field(i)
		fieldType := v.Instance.Type().Field(i)

		switch {
		case strings.Contains(fieldType.Name, "IsExecutable"):
			if err := v.setIsExecutableByMethod(process, field.Bool()); err != nil {
				return err
			}
		case fieldType.Name == "Process":
			if err := v.setProcessIdentifiers(process, fieldType.Name, field); err != nil {
				return err
			}
		}

	}

	return nil

}

// applyProcessMethods applies all BPMN element methods to the process
func (v *ReflectValue) applyProcessMethods(processIdx int, q *Quantity) error {
	methods := GetProcessMethods(processIdx, q)
	for _, methodCall := range methods {
		if methodCall.Arg <= 0 {
			continue
		}
		if err := callMethod(v.Process[processIdx], methodCall.Name,
			[]reflect.Value{reflect.ValueOf(methodCall.Arg)}); err != nil {
			return fmt.Errorf("failed to apply method %s: %w", methodCall.Name, err)
		}
	}
	return nil
}

/*
 * @ multiple processes
 */

// anonym sets the fields in the BPMN model.
func (v *ReflectValue) anonym(field string) {
	targetField := targetFieldName(v, field) // must be a struct, which represents a process
	targetNum := targetField.NumField()      // get the number of fields in the struct

	for idx := 0; idx < targetNum; idx++ {
		name := targetField.Type().Field(idx).Name

		switch targetField.Field(idx).Kind() {

		case reflect.Bool:
			v.config(name, idx, targetField)

		case reflect.Struct:
			v.currentHash(idx, targetField)
			v.nextHash(idx, targetField)
		}

	}
}

// config sets the IsExecutable field to true if the name contains "IsExecutable" and index is 0.
// I called it config, because it is a configuration of the field.
// Note: the method is used in a multiple process and it's a third option call for the field IsExecutable.
func (v *ReflectValue) config(name string, idx int, target reflect.Value) {
	if strings.Contains(name, "IsExecutable") && idx == 0 {
		target.Field(0).SetBool(true)
	}
}

// currentHash sets the hash value of the current field if it is empty.
func (v *ReflectValue) currentHash(idx int, target reflect.Value) {
	h := fmt.Sprintf("%s", target.Field(idx).FieldByName("Hash"))
	if h == "" {
		typ := typ(target.Type().Field(idx).Name)
		hash, _ := hash(typ)
		target.Field(idx).Set(reflect.ValueOf(hash))
	}
}

// nextHash sets the hash value of the next field.
// It sets the values for typ and hash in the blueprint v.Instance.
func (v *ReflectValue) nextHash(idx int, target reflect.Value) {
	if idx+1 < target.NumField() {
		typ := typ(target.Type().Field(idx + 1).Name)
		hash, _ := hash(typ)
		target.Field(idx + 1).Set(reflect.ValueOf(hash))
	}
}

// CollaborationConfig holds configuration for collaboration setup
type CollaborationConfig struct {
	Type string
	Hash string
}

// collaboration sets up the BPMN collaboration and its participants
func (v *ReflectValue) collaboration(q *Quantity) error {
	if !q.hasParticipant() {
		return nil
	}

	if err := v.setupCollaboration(); err != nil {
		return fmt.Errorf("failed to setup collaboration: %w", err)
	}

	collaboration, err := v.getCollaboration()
	if err != nil {
		return fmt.Errorf("failed to get collaboration: %w", err)
	}

	if err := v.configureCollaboration(collaboration, q.Participant); err != nil {
		return fmt.Errorf("failed to configure collaboration: %w", err)
	}

	return nil

}

// setupCollaboration initializes the collaboration
func (v *ReflectValue) setupCollaboration() error {
	method := v.Def.MethodByName("SetCollaboration")
	if !method.IsValid() {
		return fmt.Errorf("SetCollaboration method not found")
	}

	method.Call([]reflect.Value{})

	return nil

}

// getCollaboration retrieves the collaboration object
func (v *ReflectValue) getCollaboration() (reflect.Value, error) {
	method := v.Def.MethodByName("GetCollaboration")
	if !method.IsValid() {
		return reflect.Value{}, fmt.Errorf("GetCollaboration method not found")
	}

	result := method.Call([]reflect.Value{})
	if len(result) == 0 {
		return reflect.Value{}, fmt.Errorf("GetCollaboration returned no value")
	}

	return result[0], nil

}

// configureCollaboration sets up the collaboration with its participants
func (v *ReflectValue) configureCollaboration(collaboration reflect.Value, participantCount int) error {

	// get the hash value of the collaboration to add it to the ID
	collab := v.Pool.FieldByName("Collaboration").Interface().(BPMN)
	config := CollaborationConfig{
		Type: "collaboration", // Note: reconsider the value
		Hash: collab.Hash,
	}

	if err := v.setCollaborationProperties(collaboration, config, participantCount); err != nil {
		return err
	}

	return v.setupParticipants(collaboration, participantCount)

}

// setCollaborationProperties configures basic collaboration properties
func (v *ReflectValue) setCollaborationProperties(collaboration reflect.Value, config CollaborationConfig, participantCount int) error {

	if err := callMethod(collaboration, "SetID",
		[]reflect.Value{
			reflect.ValueOf(config.Type),
			reflect.ValueOf(config.Hash),
		}); err != nil {
		return fmt.Errorf("failed to set collaboration ID: %w", err)
	}

	// Set participant count
	if err := callMethod(collaboration, "SetParticipant",
		[]reflect.Value{reflect.ValueOf(participantCount)}); err != nil {
		return fmt.Errorf("failed to set participant count: %w", err)
	}

	return nil

}

// ParticipantDetails contains participant-specific information
type ParticipantDetails struct {
	ID          string
	Hash        string
	Name        string
	ProcessRef  string
	ProcessHash string
}

// setupParticipants configures all participants in the collaboration
func (v *ReflectValue) setupParticipants(collaboration reflect.Value, participantCount int) error {

	for idx := 0; idx < participantCount; idx++ {

		participant, err := v.getParticipant(collaboration, idx)
		if err != nil {
			return fmt.Errorf("failed to get participant %d: %w", idx, err)
		}

		details := ParticipantDetails{
			ID:          "participant",
			Hash:        v.ParticipantHash[idx],
			Name:        v.ParticipantName[idx],
			ProcessRef:  "process",
			ProcessHash: v.ProcessHash[idx],
		}

		if err := v.configureParticipant(participant, details); err != nil {
			return fmt.Errorf("failed to configure participant %d: %w", idx, err)
		}

	}

	return nil

}

// getParticipant retrieves a specific participant by index
func (v *ReflectValue) getParticipant(collaboration reflect.Value, index int) (reflect.Value, error) {
	result := collaboration.MethodByName("GetParticipant").Call([]reflect.Value{reflect.ValueOf(index)})
	if len(result) == 0 {
		return reflect.Value{}, fmt.Errorf("failed to get participant at index %d", index)
	}
	return result[0], nil
}

// configureParticipant sets up an individual participant
func (v *ReflectValue) configureParticipant(participant reflect.Value, details ParticipantDetails) error {
	if !participant.IsValid() {
		return fmt.Errorf("invalid participant value")
	}
	if details.Hash == "" || details.Name == "" {
		return fmt.Errorf("invalid participant details")
	}
	methods := map[string][]reflect.Value{
		"SetID":         {reflect.ValueOf(details.ID), reflect.ValueOf(details.Hash)},
		"SetName":       {reflect.ValueOf(details.Name)},
		"SetProcessRef": {reflect.ValueOf(details.ProcessRef), reflect.ValueOf(details.ProcessHash)},
	}
	for methodName, args := range methods {
		if err := callMethod(participant, methodName, args); err != nil {
			return fmt.Errorf("failed to call %s: %w", methodName, err)
		}
	}
	return nil
}

// multipleProcess processes multiple processes in the BPMN model.
func (v *ReflectValue) multipleProcess(q *Quantity) error {
	q.RLock()
	defer q.RUnlock()
	if err := v.configurePool(); err != nil {
		return fmt.Errorf("failed to configure multiple processes: %w", err)
	}
	for processIdx := 0; processIdx < q.Process; processIdx++ {

		process := v.Process[processIdx]
		isExecutable := v.ProcessExec[processIdx]
		typ := v.ProcessType[processIdx]
		hash := v.ProcessHash[processIdx]
		name := v.ProcessName[processIdx]
		elements := q.Elements[processIdx]

		if processIdx == 0 {
			if err := v.setIsExecutableByMethod(process, isExecutable); err != nil {
				return fmt.Errorf("failed to set process executable: %w", err)
			}
		}

		// set the process id
		if err := callMethod(process, "SetID",
			[]reflect.Value{reflect.ValueOf(typ), reflect.ValueOf(hash)}); err != nil {
			return fmt.Errorf("failed to apply method %s: %w", "SetID", err)
		}

		// set the process name
		if err := callMethod(process, "SetName",
			[]reflect.Value{reflect.ValueOf(name)}); err != nil {
			return fmt.Errorf("failed to apply method %s: %w", "SetName", err)
		}

		// set the process elements
		for method, arg := range elements {
			methodName := fmt.Sprintf("Set%s", method)
			if arg > 0 {
				if err := callMethod(process, methodName,
					[]reflect.Value{reflect.ValueOf(arg)}); err != nil {
					return fmt.Errorf("failed to apply method %s: %w", methodName, err)
				}
			}
		}

	}

	return nil

}

// configurePool configures multiple BPMN processes, which are in pool data structure
func (v *ReflectValue) configurePool() error {
	if v.Pool.NumField() == 0 {
		return fmt.Errorf("invalid pool: Pool is empty")
	}

	l, j, n := 0, 0, 0
	for i := 0; i < v.Pool.NumField(); i++ {
		name := v.Pool.Type().Field(i).Name
		field := v.Pool.Field(i)

		switch {
		case strings.Contains(name, "IsExecutable"):
			if field.IsValid() {
				v.ProcessExec[j] = field.Bool()
				j++
			}
		case strings.Contains(name, "Process"):
			if field.IsValid() {
				v.ProcessHash[l] = field.FieldByName("Hash").String()
				v.ProcessType[l] = field.FieldByName("Type").String()
				l++
			}
		case strings.Contains(name, "Participant"):
			if field.IsValid() {
				v.ParticipantHash[n] = field.FieldByName("Hash").String()
				n++
			}
		}

	}

	return nil

}

// cleanup clears the ReflectValue slices
// Note: memory management
func (v *ReflectValue) Cleanup() {
	// Clear slices
	v.Process = nil
	v.ProcessName = nil
	v.ProcessType = nil
	v.ProcessHash = nil
	v.ProcessExec = nil
	v.ParticipantHash = nil
	v.ParticipantName = nil
}
