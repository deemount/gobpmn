package time

import "fmt"

// TimeCycle ...
type TimeCycleRepository interface{ TimeBase }

// TimeCycle ...
type TimeCycle struct {
	TimerDefType string `xml:"xsi:type,attr,omitempty" json:"timerDefType,omitempty"`
	TimerDef     string `xml:",innerxml" json:"timerDef,omitempty"`
}

// NewTimeCycle ...
func NewTimeCycle() TimeCycleRepository {
	return &TimeCycle{}
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
func (timecycle TimeCycle) GetTimerDefinitionType() STR_PTR {
	return &timecycle.TimerDefType
}

// GetTimerDefinition ...
func (timecycle TimeCycle) GetTimerDefinition() STR_PTR {
	return &timecycle.TimerDef
}
