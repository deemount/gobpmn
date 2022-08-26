package models

import "fmt"

// SequenceFlowRepository ...
type SequenceFlowRepository interface {
	SetID(suffix string)
	SetName(name string)
}

// SequenceFlow ...
type SequenceFlow struct {
	ID                  string                `xml:"id,attr" json:"id"`
	Name                string                `xml:"name,attr,omitempty" json:"name,omitempty"`
	SourceRef           string                `xml:"sourceRef,attr" json:"sourceRef"`
	TargetRef           string                `xml:"targetRef,attr" json:"targetRef"`
	ConditionExpression []ConditionExpression `xml:"bpmn:conditionExpression,omitempty" json:"conditionExpression,omitempty"`
}

// TSequenceFlow ...
type TSequenceFlow struct {
	ID                  string                `xml:"id,attr" json:"id"`
	Name                string                `xml:"name,attr,omitempty" json:"name,omitempty"`
	SourceRef           string                `xml:"sourceRef,attr" json:"sourceRef"`
	TargetRef           string                `xml:"targetRef,attr" json:"targetRef"`
	ConditionExpression []ConditionExpression `xml:"conditionExpression,omitempty" json:"conditionExpression,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (sequenceFlow *SequenceFlow) SetID(suffix string) {
	sequenceFlow.ID = fmt.Sprintf("Flow_%s", suffix)
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
	case "startevent":
		sequenceFlow.SourceRef = fmt.Sprintf("StartEvent_%s", sourceRef)
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

/* Elements */

/** BPMN **/

// SetConditionExpression ...
func (sequenceFlow *SequenceFlow) SetConditionExpression() {
	sequenceFlow.ConditionExpression = make([]ConditionExpression, 1)
}
