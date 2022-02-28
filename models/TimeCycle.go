package models

import "fmt"

// TimeCycle ...
type TimeCycle struct {
	TimerDefType string `xml:"xsi:type,attr,omitempty" json:"timerDefType,omitempty"`
	TimerDef     string `xml:",innerxml" json:"timerDef,omitempty"`
}

/* Attributes */

// SetTimerDefinitionType ...
func (timecycle *TimeCycle) SetTimerDefinitionType() {
	timecycle.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timecycle *TimeCycle) SetTimerDefinition(timerDefinition string) {
	timecycle.TimerDef = timerDefinition
}
