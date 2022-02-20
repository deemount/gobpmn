package models

import "fmt"

// TimeDate ...
type TimeDate struct {
	TimerDefType string `xml:"xsi:type,attr,omitempty"`
	TimerDef     string `xml:",innerxml"`
}

// SetTimerDefinitionType ...
func (timedate *TimeDate) SetTimerDefinitionType() {
	timedate.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timedate *TimeDate) SetTimerDefinition(timerDefinition string) {
	timedate.TimerDef = timerDefinition
}
