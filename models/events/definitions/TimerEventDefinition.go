package definitions

import (
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/time"
)

// NewTimerEventDefinition ...
func NewTimerEventDefinition() TimerEventDefinitionRepository {
	return &TimerEventDefinition{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ted *TimerEventDefinition) SetID(typ string, suffix interface{}) {
	ted.ID = SetID(typ, suffix)
}

/*** Make Elements ***/

/** BPMN **/

// SetTimeDuration ...
func (ted *TimerEventDefinition) SetTimeDate() {
	ted.TimeDate = make([]time.TimeDate, 1)
}

// SetTimeDuration ...
func (ted *TimerEventDefinition) SetTimeDuration() {
	ted.TimeDuration = make([]time.TimeDuration, 1)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ted TimerEventDefinition) GetID() compulsion.STR_PTR {
	return &ted.ID
}

/* Elements */

/** BPMN **/

// GetTimeDuration ...
func (ted TimerEventDefinition) GetTimeDate() *time.TimeDate {
	return &ted.TimeDate[0]
}

// GetTimeDuration ...
func (ted TimerEventDefinition) GetTimeDuration() *time.TimeDuration {
	return &ted.TimeDuration[0]
}
