package time

import "fmt"

// TimeDurationRepository ...
type TimeDurationRepository interface{ TimeBase }

// TimeDuration ...
type TimeDuration struct {
	TimerDefType string `xml:"xsi:type,attr,omitempty" json:"timerDefType,omitempty"`
	TimerDef     string `xml:",innerxml" json:"timerDef,omitempty"`
}

// NewTimeDuration ...
func NewTimeDuration() TimeDurationRepository {
	return &TimeDuration{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetTimerDefinitionType ...
func (timeduration *TimeDuration) SetTimerDefinitionType() {
	timeduration.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timeduration *TimeDuration) SetTimerDefinition(timerDefinition string) {
	timeduration.TimerDef = timerDefinition
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetTimerDefinitionType ...
func (timeduration TimeDuration) GetTimerDefinitionType() STR_PTR {
	return &timeduration.TimerDefType
}

// GetTimerDefinition ...
func (timeduration TimeDuration) GetTimerDefinition() STR_PTR {
	return &timeduration.TimerDef
}
