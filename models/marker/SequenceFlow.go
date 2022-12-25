package marker

import (
	"fmt"

	"github.com/deemount/gobpmn/models/conditional"
)

// SequenceFlowRepository ...
type SequenceFlowRepository interface {
	MarkerBase

	SetSourceRef(typ string, sourceRef interface{})
	SetTargetRef(typ string, targetRef interface{})
	GetSourceRef() *string
	GetTargetRef() *string

	SetConditionExpression()
	GetConditionExpression() *conditional.ConditionExpression
}

// SequenceFlow ...
type SequenceFlow struct {
	ID                  string                            `xml:"id,attr" json:"id"`
	Name                string                            `xml:"name,attr,omitempty" json:"name,omitempty"`
	SourceRef           string                            `xml:"sourceRef,attr" json:"sourceRef"`
	TargetRef           string                            `xml:"targetRef,attr" json:"targetRef"`
	ConditionExpression []conditional.ConditionExpression `xml:"bpmn:conditionExpression,omitempty" json:"conditionExpression,omitempty"`
}

// TSequenceFlow ...
type TSequenceFlow struct {
	ID                  string                            `xml:"id,attr" json:"id"`
	Name                string                            `xml:"name,attr,omitempty" json:"name,omitempty"`
	SourceRef           string                            `xml:"sourceRef,attr" json:"sourceRef"`
	TargetRef           string                            `xml:"targetRef,attr" json:"targetRef"`
	ConditionExpression []conditional.ConditionExpression `xml:"conditionExpression,omitempty" json:"conditionExpression,omitempty"`
}

func NewSequenceFlow() SequenceFlowRepository {
	return &SequenceFlow{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (sequenceFlow *SequenceFlow) SetID(typ string, suffix interface{}) {
	switch typ {
	case "flow":
		sequenceFlow.ID = fmt.Sprintf("Flow_%v", suffix)
		break
	case "id":
		sequenceFlow.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (sequenceFlow *SequenceFlow) SetName(name string) {
	sequenceFlow.Name = name
}

// SetSourceRef ...
func (sequenceFlow *SequenceFlow) SetSourceRef(typ string, sourceRef interface{}) {
	switch typ {
	case "activity":
		sequenceFlow.SourceRef = fmt.Sprintf("Activity_%s", sourceRef)
		break
	case "event":
		sequenceFlow.SourceRef = fmt.Sprintf("Event_%s", sourceRef)
		break
	// Notice: using %v instead %s for more flexible parameter values
	case "startevent":
		sequenceFlow.SourceRef = fmt.Sprintf("StartEvent_%v", sourceRef)
		break
	case "id":
		sequenceFlow.SourceRef = fmt.Sprintf("%s", sourceRef)
		break
	}
}

// SetTargetRef ...
func (sequenceFlow *SequenceFlow) SetTargetRef(typ string, targetRef interface{}) {
	switch typ {
	case "activity":
		sequenceFlow.TargetRef = fmt.Sprintf("Activity_%s", targetRef)
		break
	case "event":
		sequenceFlow.TargetRef = fmt.Sprintf("Event_%s", targetRef)
		break
	case "startevent":
		sequenceFlow.TargetRef = fmt.Sprintf("StartEvent_%s", targetRef)
		break
	case "id":
		sequenceFlow.TargetRef = fmt.Sprintf("%s", targetRef)
		break
	}
}

/*** Make Elements ***/

/** BPMN **/

// SetConditionExpression ...
func (sequenceFlow *SequenceFlow) SetConditionExpression() {
	sequenceFlow.ConditionExpression = make([]conditional.ConditionExpression, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (sequenceFlow SequenceFlow) GetID() STR_PTR {
	return &sequenceFlow.ID
}

// GetName ...
func (sequenceFlow SequenceFlow) GetName() STR_PTR {
	return &sequenceFlow.Name
}

// GetSourceRef ...
func (sequenceFlow SequenceFlow) GetSourceRef() *string {
	return &sequenceFlow.SourceRef
}

// GetTargetRef ...
func (sequenceFlow SequenceFlow) GetTargetRef() *string {
	return &sequenceFlow.TargetRef
}

/* Elements */

/** BPMN **/

// GetConditionExpression ...
func (sequenceFlow SequenceFlow) GetConditionExpression() *conditional.ConditionExpression {
	return &sequenceFlow.ConditionExpression[0]
}