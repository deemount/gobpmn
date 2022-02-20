package models

import "fmt"

// TimeDuration ...
type TimeDuration struct {
	TimerDefType string `xml:"xsi:type,attr,omitempty"`
	TimerDef     string `xml:",innerxml"`
}

// SetTimerDefinitionType ...
func (timeduration *TimeDuration) SetTimerDefinitionType() {
	timeduration.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timeduration *TimeDuration) SetTimerDefinition(timerDefinition string) {
	timeduration.TimerDef = timerDefinition
}
