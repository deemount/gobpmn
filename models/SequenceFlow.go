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
func (sequenceFlow *SequenceFlow) SetSourceRef(sourceRef string) {
	sequenceFlow.SourceRef = sourceRef
}

// SetTargetRef ...
func (sequenceFlow *SequenceFlow) SetTargetRef(targetRef string) {
	sequenceFlow.TargetRef = targetRef
}

/* Elements */

/** BPMN **/

// SetConditionExpression ...
func (sequenceFlow *SequenceFlow) SetConditionExpression() {
	sequenceFlow.ConditionExpression = make([]ConditionExpression, 1)
}
