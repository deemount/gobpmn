package common

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/deemount/gobpmn/pkg/config"
	"github.com/deemount/gobpmn/pkg/core/foundation"
)

// ReflectValue represents the core structure for BPMN reflection
type ReflectValue[M BPMNGeneric] struct {
	config.ModelConfig[M]
	config.ProcessConfig
	config.ParticipantConfig
	FlowNeighbors    map[string]map[string]any
	InstanceNumField int
}

// NewReflectValue is the first call in the core.ReflectDI function.
// It creates a new instance of the ReflectValue and saves it in ModelConfig.
func NewReflectValue[T any, M BPMNGeneric](model T, genericType M) (*ReflectValue[M], error) {
	if any(model) == nil {
		return nil, NewError(fmt.Errorf("model cannot be nil"))
	}
	var visibleFields M
	var instance reflect.Value
	// reflect the type of the struct (typeOf) and
	typeOf := reflect.TypeOf(model)
	valueOf := reflect.ValueOf(model)
	switch typeOf.Kind() {
	case reflect.Struct:
		fields := reflect.VisibleFields(typeOf)
		visibleFields = any(fields).(M)
		instance = reflect.New(typeOf).Elem()
		return &ReflectValue[M]{
			ModelConfig: config.ModelConfig[M]{
				Name:     extractPrefixBeforeProcess(typeOf.Name()),
				Type:     typeOf,
				Wrap:     valueOf,
				Fields:   visibleFields, // struct fields or map keys
				Instance: instance,      // create a new object of the struct/map, which is a copy of the model
			},
		}, nil
	case reflect.Map:
		fields := make(map[string]any)
		var sortedKeys []string
		// temporary struct to store BPMN elements
		bpmnElements := make([]struct {
			Key string
			Pos int
		}, 0)
		// iterate over the map keys
		for _, key := range valueOf.MapKeys() {
			strKey := key.String()
			value := valueOf.MapIndex(key).Interface()
			switch v := value.(type) {
			case BPMN:
				bpmnElements = append(bpmnElements, struct {
					Key string
					Pos int
				}{strKey, v.Pos})
				fields[strKey] = v
			case bool, any:
				fields[strKey] = v
			}
		}
		// sort BPMN elements by Pos
		sort.SliceStable(bpmnElements, func(i, j int) bool {
			return bpmnElements[i].Pos < bpmnElements[j].Pos
		})
		// create sorted keys
		for _, elem := range bpmnElements {
			sortedKeys = append(sortedKeys, elem.Key)
		}
		visibleFields := fields // type assertion

		converter := &Converter{}
		result, err := converter.ToStruct(model)
		if err != nil {
			return nil, NewError(fmt.Errorf("%v\n", err))
		}
		instance = reflect.New(reflect.TypeOf(result.Value.Interface())).Elem()
		return &ReflectValue[M]{
			ModelConfig: config.ModelConfig[M]{
				Name:     extractPrefixBeforeProcess(typeOf.Name()),
				Type:     typeOf,
				Wrap:     valueOf,
				Fields:   any(visibleFields).(M), // struct fields or map keys
				Instance: instance,               // create a new object of the struct/map, which is a copy of the model
			},
		}, nil
	}
	return nil, NewError(fmt.Errorf("unsupported type: %s", typeOf.Kind()))

}

// Setup sets up the initial configuration for the ReflectValue
func (v *ReflectValue[M]) Setup() error {
	if err := v.setupInstanceFields(); err != nil {
		return NewError(fmt.Errorf("failed to setup target fields:\n%w", err))
	}
	if err := v.setupDefinition(); err != nil {
		return NewError(fmt.Errorf("failed to setup definition:\n%w", err))
	}
	return nil
}

// setupInstanceFields initializes the target fields
// NOTE: Actually, it setups the number of fields of the model
// and the definition field.
func (v *ReflectValue[M]) setupInstanceFields() error {
	// set the number of fields, if M is a struct
	v.InstanceNumField = v.Instance.NumField()
	v.Def = instanceFieldName(v, "Def")
	if !v.Def.IsValid() {
		return NewError(fmt.Errorf("definition field is not valid"))
	}
	return nil
}

// setupDefinition configures the definition field
func (v *ReflectValue[M]) setupDefinition() error {
	if !v.Def.CanSet() {
		return nil
	}
	definitions := foundation.NewDefinitions()
	definitions.SetDefaultAttributes()
	v.Def.Set(reflect.ValueOf(definitions))
	return nil
}

// handleModelType determines and handles the model type (pool or single)
func (v *ReflectValue[M]) HandleModelType(q *Quantity[M], m *Mapping[M]) error {
	if len(m.Anonym) > 0 {
		return v.handlePool(q, m)
	}
	return v.handleStandalone(q, m)
}

// handlePool assigns the field Pool to v.Pool.
// It walks through the anonymous fields in Mapping and checks
// if the field contains the word "Pool".
// A pool can only be set once in the BPMN model.
func (v *ReflectValue[M]) handlePool(q *Quantity[M], m *Mapping[M]) error {
	for _, field := range m.Anonym {
		if strings.Contains(field, "Pool") {
			v.Pool = instanceFieldName(v, field)
			q.countFieldsInPool(v)
			if !v.Pool.IsValid() {
				return NewError(fmt.Errorf("invalid pool field: %s", field))
			}
			q.countFieldsInInstance(v)
			q.Pool++
			break
		}
	}
	return nil
}

// handleStandalone appends the process name to v.ProcessName.
// The process name is extracted in NewReflectValue, which is
// the name of the data structure.
//
// NOTE: A standalone process is a process without a pool.
func (v *ReflectValue[M]) handleStandalone(q *Quantity[M], m *Mapping[M]) error {
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

// Instnc holds the data structure of the BPMN model in the ReflectValue.
// the return value is then as a return value in the initial function.
//
// NOTE:
// It is used as a blueprint for the BPMN model, which is then reflected into v.Def.
// It distribute hash values and set the isExecutable field to true.
// It uses the type declaration in the data structure to generate the hash value.
func (v *ReflectValue[M]) Instnc(m *Mapping[M]) any {
	if len(m.Anonym) > 0 {
		for _, anonymField := range m.Anonym {
			v.anonym(anonymField)
		}
	} else {
		v.nonAnonym(m)
	}
	// NOTE:
	// FlowNeighbors are used in the element_handler.go file,
	// for the FlowHandler Properties.
	// It works for multiple and/or standalone models.
	v.FlowNeighbors = collectFromFieldsWithNeighbors(v.Instance.Interface())
	return v.Instance.Interface()
}

// nonAnonym sets the IsExecutable field to true and populates the reflection fields with hash values.
// NOTE: the method is used in a standalone process.
func (v *ReflectValue[M]) nonAnonym(m *Mapping[M]) {
	switch v.Instance.Kind() {
	case reflect.Struct:
		v.setIsExecutableByField(m.Config)
		v.populateReflectionFields(m.BPMNType)
	case reflect.Map:
		v.populateReflectionFields(m.BPMNType)
	}
}

// anonym sets the fields in the BPMN model.
func (v *ReflectValue[M]) anonym(field string) {
	targetField := instanceFieldName(v, field) // must be a struct, which represents a process
	targetNum := targetField.NumField()        // get the number of fields in the struct
	for idx := range targetNum {
		name := targetField.Type().Field(idx).Name
		switch targetField.Field(idx).Kind() {
		case reflect.Bool:
			v.config(name, idx, targetField)
		case reflect.Struct:
			// NOTE: need to have the option, which process is it.
			// I need it for the first part of the bpmn injection:
			// e.g.: StartEvent_HashValue -> Event_HashValue.
			v.currentHash(idx, targetField)
			v.nextHash(idx, targetField)
		}
	}
}

// ReflectProcess reflects the processes in the BPMN model.
// v.Process[index of process] is a slice of reflect.Value, which represents the processes in the BPMN model.
// v.Process[0] is the first process in the model and it's used also in a standalone process arrangement.
func (v *ReflectValue[M]) ReflectProcess(q *Quantity[M]) error {
	v.Process = make([]reflect.Value, q.Process)
	// extract the struct, if v.Def is a struct
	def := v.Def
	if def.Kind() == reflect.Interface || def.Kind() == reflect.Ptr {
		def = reflect.ValueOf(def.Interface()) // get the real value by any
	}
	method := def.MethodByName("SetProcess")
	if !method.IsValid() {
		return NewError(fmt.Errorf("SetProcess method not found"))
	}
	results := method.Call([]reflect.Value{reflect.ValueOf(q.Process)})
	if len(results) > 0 && !results[0].IsNil() {
		return NewError(fmt.Errorf("SetProcess call failed"))
	}
	for processIdx := range q.Process {
		v.Process[processIdx] = def.MethodByName("GetProcess").Call([]reflect.Value{reflect.ValueOf(processIdx)})[0]
	}
	return nil
}

// processingWithContext processes the BPMN model with a cancellable context.
func (v *ReflectValue[M]) processingWithContext(ctx context.Context, q *Quantity[M]) error {
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

// processing sets the maps and process parameters in the BPMN model.
// If the pool is greater than 0, then process multiple processes.
// It considers the pool as a collaboration.
// It creates a cancellable context for the processor.
func (v *ReflectValue[M]) processing(ctx context.Context, q *Quantity[M]) error {
	// initialize the slices
	v.ProcessExec = make([]bool, q.Process)
	v.ProcessType = make([]string, q.Process)
	v.ProcessHash = make([]string, q.Process)
	v.ParticipantHash = make([]string, q.Participant)
	// create a new element processor
	processor, err := NewElementProcessor(v, q)
	if err != nil {
		return NewError(fmt.Errorf("failed to create element processor:\n%w", err))
	}
	// create a cancellable context
	processorCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	// a pool is detected
	if q.Pool > 0 {
		err := v.configurePool()
		if err != nil {
			return NewError(fmt.Errorf("error configuring pool:\n%v", err))
		}
		v.multipleProcess(q)
		if err := processor.ProcessElementsWithContext(processorCtx); err != nil {
			return err
		}
	} else {
		if err := v.configureStandalone(); err != nil {
			return NewError(fmt.Errorf("error processing standalone:\n%v", err))
		}
		v.standalone(0, q)
		if err := processor.ProcessStandalone(processorCtx); err != nil {
			return err
		}
	}
	return nil
}

// FinalizeModel completes the model construction
func (v *ReflectValue[M]) FinalizeModel(ctx context.Context, q *Quantity[M]) error {
	if err := v.processingWithContext(ctx, q); err != nil {
		return NewError(fmt.Errorf("failed to process:\n%w", err))
	}
	if err := v.collaboration(q); err != nil {
		return NewError(fmt.Errorf("failed to setup collaboration:\n%w", err))
	}
	return nil
}

// setIsExecutableByField configures the process executable status.
// It sets the IsExecutable field to true.
// NOTE: maybe redudant; look at method config.
// NOTE: this method makes no really sense, if M is map[string]any. IsExecutable can be set directly in the map.
func (v *ReflectValue[M]) setIsExecutableByField(fields map[int]string) {
	for _, field := range fields {
		instanceFieldName(v, field).SetBool(true)
		break
	}
}

// setIsExecutableByMethod configures the process executable status.
// It calls SetIsExecutable method on the process.
func (v *ReflectValue[M]) setIsExecutableByMethod(process reflect.Value, isExecutable bool) error {
	method := process.MethodByName("SetIsExecutable")
	if !method.IsValid() {
		return NewError(fmt.Errorf("SetIsExecutable method not found"))
	}
	method.Call([]reflect.Value{reflect.ValueOf(isExecutable)})
	return nil
}

// populateReflectionFields populates the reflection fields with hash values.
// It sets the hash values for the BPMN elements in the model.
// NOTE: the method is used in a standalone process.
func (v *ReflectValue[M]) populateReflectionFields(reflectionFields map[int]string) {
	for _, field := range reflectionFields {
		f := instanceFieldName(v, field)
		if !f.IsValid() {
			fmt.Printf("populateReflectionFields: field %s not found in struct\n", field)
			continue
		}
		fName, _ := v.Instance.Type().FieldByName(field)
		typ := typ(fName.Name)
		hash, _ := inject(typ, 0)
		f.Set(reflect.ValueOf(hash))
	}
}

// setProcessIdentifiers sets the process ID and name
func (v *ReflectValue[M]) setProcessIdentifiers(process reflect.Value, name string, field reflect.Value) error {
	var hash string
	hash = field.FieldByName("Hash").String()
	// set process id
	if err := callMethod(process, "SetID",
		[]reflect.Value{
			reflect.ValueOf(strings.ToLower(name)),
			reflect.ValueOf(hash),
		}); err != nil {
		return NewError(fmt.Errorf("failed to set process ID:\n%w", err))
	}
	// set process name
	if err := callMethod(process, "SetName",
		[]reflect.Value{reflect.ValueOf(name)}); err != nil {
		return NewError(fmt.Errorf("failed to set process name:\n%w", err))
	}
	return nil
}

// standalone configures and initializes a single BPMN process.
// It sets up the process properties and applies the required BPMN elements.
// NOTE: this method is used for a single process and in usage in the method "process".
func (v *ReflectValue[M]) standalone(processIdx int, q *Quantity[M]) error {
	if err := v.configureStandalone(); err != nil {
		return NewError(fmt.Errorf("failed to configure process:\n%w", err))
	}
	if err := v.applyProcessMethods(processIdx, q); err != nil {
		return NewError(fmt.Errorf("failed to apply process methods:\n%w", err))
	}
	return nil
}

// configureStandalone configures a single BPMN process.
func (v *ReflectValue[M]) configureStandalone() error {
	if v.Process == nil || len(v.Process) == 0 {
		return fmt.Errorf("invalid process: Process slice is nil or empty")
	}
	process := v.Process[0] // get the first process
	// processing for structs
	for i := range v.InstanceNumField {
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
func (v *ReflectValue[M]) applyProcessMethods(processIdx int, q *Quantity[M]) error {
	methods := GetProcessMethods(processIdx, q)
	for _, methodCall := range methods {
		if methodCall.Arg <= 0 {
			continue
		}
		if err := callMethod(v.Process[processIdx], methodCall.Name,
			[]reflect.Value{reflect.ValueOf(methodCall.Arg)}); err != nil {
			return NewError(fmt.Errorf("failed to apply method %s:\n%w", methodCall.Name, err))
		}
	}
	return nil
}

/*
 * @ multiple processes
 */

// config sets the IsExecutable field to true if the name contains "IsExecutable" and index is 0.
// I called it config, because it is a configuration of the field.
// Note: the method is used in a multiple process and it's a third option call for the field IsExecutable.
func (v *ReflectValue[M]) config(name string, idx int, target reflect.Value) {
	if strings.Contains(name, "IsExecutable") && idx == 0 {
		target.Field(0).SetBool(true)
	}
}

// currentHash sets the hash value of the current field if it is empty.
// It sets the values for typ and hash in the blueprint v.Instance.
// NOTE:
// The variable typ is the type of the BPMN element,
// which is token by the name of the field. NOT THE BEST PRACTICE.
func (v *ReflectValue[M]) currentHash(idx int, target reflect.Value) {
	h := fmt.Sprintf("%s", target.Field(idx).FieldByName("Hash"))
	if h == "" {
		typ := typ(target.Type().Field(idx).Name)
		hash, _ := inject(typ, idx)
		target.Field(idx).Set(reflect.ValueOf(hash))
	}
}

// nextHash sets the hash value of the next field.
// It sets the values for typ and hash in the blueprint v.Instance.
func (v *ReflectValue[M]) nextHash(idx int, target reflect.Value) {
	if idx+1 < target.NumField() {
		typ := typ(target.Type().Field(idx + 1).Name)
		hash, _ := inject(typ, idx)
		target.Field(idx + 1).Set(reflect.ValueOf(hash))
	}
}

// CollaborationConfig holds configuration for collaboration setup

// collaboration sets up the BPMN collaboration and its participants
func (v *ReflectValue[M]) collaboration(q *Quantity[M]) error {
	if !q.hasParticipant() {
		return nil
	}
	if err := v.setupCollaboration(); err != nil {
		return NewError(fmt.Errorf("failed to setup collaboration: %w", err))
	}
	collaboration, err := v.getCollaboration()
	if err != nil {
		return NewError(fmt.Errorf("failed to get collaboration: %w", err))
	}
	if err := v.configureCollaboration(collaboration, q.Participant); err != nil {
		return NewError(fmt.Errorf("failed to configure collaboration: %w", err))
	}
	return nil
}

// setupCollaboration initializes the collaboration
func (v *ReflectValue[M]) setupCollaboration() error {
	method := v.Def.MethodByName("SetCollaboration")
	if !method.IsValid() {
		return NewError(fmt.Errorf("SetCollaboration method not found"))
	}
	method.Call([]reflect.Value{})
	return nil
}

// getCollaboration retrieves the collaboration object
func (v *ReflectValue[M]) getCollaboration() (reflect.Value, error) {
	method := v.Def.MethodByName("GetCollaboration")
	if !method.IsValid() {
		return reflect.Value{}, NewError(fmt.Errorf("GetCollaboration method not found"))
	}
	result := method.Call([]reflect.Value{})
	if len(result) == 0 {
		return reflect.Value{}, NewError(fmt.Errorf("GetCollaboration returned no value"))
	}
	return result[0], nil
}

// configureCollaboration sets up the collaboration with its participants
func (v *ReflectValue[M]) configureCollaboration(collaboration reflect.Value, participantCount int) error {
	// get the hash value of the collaboration to add it to the ID
	collab := v.Pool.FieldByName("Collaboration").Interface().(BPMN)
	config := config.CollaborationConfig{
		Type: "collaboration", // Note: reconsider the value
		Hash: collab.Hash,
	}
	if err := v.setCollaborationProperties(collaboration, config, participantCount); err != nil {
		return err
	}
	return v.setupParticipants(collaboration, participantCount)
}

// setCollaborationProperties configures basic collaboration properties
func (v *ReflectValue[M]) setCollaborationProperties(collaboration reflect.Value, config config.CollaborationConfig, participantCount int) error {
	if err := callMethod(collaboration, "SetID",
		[]reflect.Value{
			reflect.ValueOf(config.Type),
			reflect.ValueOf(config.Hash),
		}); err != nil {
		return NewError(fmt.Errorf("failed to set collaboration ID: %w", err))
	}
	// Set participant count
	if err := callMethod(collaboration, "SetParticipant",
		[]reflect.Value{reflect.ValueOf(participantCount)}); err != nil {
		return NewError(fmt.Errorf("failed to set participant count: %w", err))
	}
	return nil
}

// setupParticipants configures all participants in the collaboration
func (v *ReflectValue[M]) setupParticipants(collaboration reflect.Value, participantCount int) error {
	for idx := range participantCount {
		participant, err := v.getParticipant(collaboration, idx)
		if err != nil {
			return NewError(fmt.Errorf("failed to get participant %d: %w", idx, err))
		}
		details := config.ParticipantDetails{
			ID:          "participant",
			Hash:        v.ParticipantHash[idx],
			Name:        v.ParticipantName[idx],
			ProcessRef:  "process",
			ProcessHash: v.ProcessHash[idx],
		}
		if err := v.configureParticipant(participant, details); err != nil {
			return NewError(fmt.Errorf("failed to configure participant %d: %w", idx, err))
		}
	}
	return nil
}

// getParticipant retrieves a specific participant by index
func (v *ReflectValue[M]) getParticipant(collaboration reflect.Value, index int) (reflect.Value, error) {
	result := collaboration.MethodByName("GetParticipant").Call([]reflect.Value{reflect.ValueOf(index)})
	if len(result) == 0 {
		return reflect.Value{}, NewError(fmt.Errorf("failed to get participant at index %d", index))
	}
	return result[0], nil
}

// configureParticipant sets up an individual participant
func (v *ReflectValue[M]) configureParticipant(participant reflect.Value, details config.ParticipantDetails) error {
	if !participant.IsValid() {
		return NewError(fmt.Errorf("invalid participant value"))
	}
	if details.Hash == "" || details.Name == "" {
		return NewError(fmt.Errorf("invalid participant details"))
	}
	methods := map[string][]reflect.Value{
		"SetID":         {reflect.ValueOf(details.ID), reflect.ValueOf(details.Hash)},
		"SetName":       {reflect.ValueOf(details.Name)},
		"SetProcessRef": {reflect.ValueOf(details.ProcessRef), reflect.ValueOf(details.ProcessHash)},
	}
	for methodName, args := range methods {
		if err := callMethod(participant, methodName, args); err != nil {
			return NewError(fmt.Errorf("failed to call %s: %w", methodName, err))
		}
	}
	return nil
}

// multipleProcess processes multiple processes in the BPMN model.
func (v *ReflectValue[M]) multipleProcess(q *Quantity[M]) error {
	q.RLock()
	defer q.RUnlock()
	if err := v.configurePool(); err != nil {
		return fmt.Errorf("failed to configure multiple processes:\n%w", err)
	}
	for processIdx := range q.Process {
		process := v.Process[processIdx]
		isExecutable := v.ProcessExec[processIdx]
		typ := v.ProcessType[processIdx]
		hash := v.ProcessHash[processIdx]
		name := v.ProcessName[processIdx]
		elements := q.Elements[processIdx]
		if processIdx == 0 {
			if err := v.setIsExecutableByMethod(process, isExecutable); err != nil {
				return NewError(fmt.Errorf("failed to set process executable:\n%w", err))
			}
		}
		// set the process id
		if err := callMethod(process, "SetID",
			[]reflect.Value{reflect.ValueOf(typ), reflect.ValueOf(hash)}); err != nil {
			return NewError(fmt.Errorf("failed to apply method %s: %w", "SetID", err))
		}
		// set the process name
		if err := callMethod(process, "SetName",
			[]reflect.Value{reflect.ValueOf(name)}); err != nil {
			return NewError(fmt.Errorf("failed to apply method %s: %w", "SetName", err))
		}
		// set the process elements
		for method, arg := range elements {
			methodName := fmt.Sprintf("Set%s", method)
			if arg > 0 {
				if err := callMethod(process, methodName,
					[]reflect.Value{reflect.ValueOf(arg)}); err != nil {
					return NewError(fmt.Errorf("failed to apply method %s: %w", methodName, err))
				}
			}
		}
	}
	return nil
}

// configurePool configures multiple BPMN processes, which are in pool data structure
func (v *ReflectValue[M]) configurePool() error {
	if v.Pool.NumField() == 0 {
		return NewError(fmt.Errorf("invalid pool: Pool is empty"))
	}
	l, j, n := 0, 0, 0
	for i := range v.Pool.NumField() {
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

// instanceFieldName returns the field value by the name in v.Instance.
func instanceFieldName[M BPMNGeneric](v *ReflectValue[M], name string) reflect.Value {
	value := v.Instance.FieldByName(name)
	if !value.IsValid() {
		fmt.Printf("instanceFieldName: field %s not found in struct instance. Available fields: %v\n", name, getFieldNames(v.Instance))
		return reflect.Value{}
	}
	return value
}

// getFieldNames returns the names of all fields in a reflect.Value instance
func getFieldNames(instance reflect.Value) []string {
	var fieldNames []string
	for i := range instance.NumField() {
		fieldNames = append(fieldNames, instance.Type().Field(i).Name)
	}
	return fieldNames
}

// Cleanup clears the ReflectValue slices
// Note: memory management
func (v *ReflectValue[M]) Cleanup() {
	// Clear slices
	v.Process = nil
	v.ProcessName = nil
	v.ProcessType = nil
	v.ProcessHash = nil
	v.ProcessExec = nil
	v.ParticipantHash = nil
	v.ParticipantName = nil
}
