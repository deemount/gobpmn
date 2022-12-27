package time

import "fmt"

// TimeCycle ...
type TimeCycleRepository interface{ TimeBase }

// TimeCycle ...
type TimeCycle struct {
	TimerDefType string `xml:"xsi:type,attr,omitempty" json:"timerDefType,omitempty"`
	TimerDef     string `xml:",innerxml" json:"timerDef,omitempty"`
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetTimerDefinitionType ...
func (timecycle *TimeCycle) SetTimerDefinitionType() {
	timecycle.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timecycle *TimeCycle) SetTimerDefinition(timerDefinition string) {
	timecycle.TimerDef = timerDefinition
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetTimerDefinitionType ...
func (timecycle TimeCycle) GetTimerDefinitionType() *string {
	return &timecycle.TimerDefType
}

// GetTimerDefinition ...
func (timecycle TimeCycle) GetTimerDefinition() *string {
	return &timecycle.TimerDef
}
