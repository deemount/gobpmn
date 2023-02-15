package time

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// NewTimeDate ...
func NewTimeDate() TimeDateRepository {
	return &TimeDate{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetTimerDefinitionType ...
func (timedate *TimeDate) SetTimerDefinitionType() {
	timedate.TimerDefType = fmt.Sprint("bpmn:tFormalExpression")
}

// SetTimerDefinition ...
func (timedate *TimeDate) SetTimerDefinition(timerDefinition string) {
	timedate.TimerDef = timerDefinition
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetTimerDefinitionType ...
func (timedate TimeDate) GetTimerDefinitionType() impl.STR_PTR {
	return &timedate.TimerDefType
}

// GetTimerDefinition ...
func (timedate TimeDate) GetTimerDefinition() impl.STR_PTR {
	return &timedate.TimerDef
}
