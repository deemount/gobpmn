package models

import "fmt"

// TimeCycle ...
type TimeCycle struct {
	TimerDefType string `xml:"xsi:type,attr,omitempty"`
	TimerDef     string `xml:",innerxml"`
}

// SetTimerDefinitionType ...
func (timecycle *TimeCycle) SetTimerDefinitionType() {
	timecycle.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timecycle *TimeCycle) SetTimerDefinition(timerDefinition string) {
	timecycle.TimerDef = timerDefinition
}
