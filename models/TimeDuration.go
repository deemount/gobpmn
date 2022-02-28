package models

import "fmt"

// TimeDuration ...
type TimeDuration struct {
	TimerDefType string `xml:"xsi:type,attr,omitempty" json:"timerDefType,omitempty"`
	TimerDef     string `xml:",innerxml" json:"timerDef,omitempty"`
}

/* Attributes */

// SetTimerDefinitionType ...
func (timeduration *TimeDuration) SetTimerDefinitionType() {
	timeduration.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timeduration *TimeDuration) SetTimerDefinition(timerDefinition string) {
	timeduration.TimerDef = timerDefinition
}
