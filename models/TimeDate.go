package models

import "fmt"

// TimeDate ...
type TimeDate struct {
	TimerDefType string `xml:"xsi:type,attr,omitempty" json:"timerDefType,omitempty"`
	TimerDef     string `xml:",innerxml" json:"timerDef,omitempty"`
}

/* Attributes */

// SetTimerDefinitionType ...
func (timedate *TimeDate) SetTimerDefinitionType() {
	timedate.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timedate *TimeDate) SetTimerDefinition(timerDefinition string) {
	timedate.TimerDef = timerDefinition
}
