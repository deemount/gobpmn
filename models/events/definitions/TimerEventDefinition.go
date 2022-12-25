package definitions

import (
	"fmt"

	"github.com/deemount/gobpmn/models/events/eventsbase"
	"github.com/deemount/gobpmn/models/time"
)

// TimerEventDefinitionRepository ...
type TimerEventDefinitionRepository interface {
	DefinitionsID

	SetTimeDate()
	GetTimeDate() *time.TimeDate
	SetTimeDuration()
	GetTimeDuration() *time.TimeDuration
}

// TimerEventDefinition ...
type TimerEventDefinition struct {
	ID           string              `xml:"id,attr,omitempty" json:"id"`
	TimeDate     []time.TimeDate     `xml:"bpmn:timeDate,omitempty" json:"timeDate,omitempty"`
	TimeDuration []time.TimeDuration `xml:"bpmn:timeDuration,omitempty" json:"timeDuration,omitempty"`
}

// TTimerEventDefinition ...
type TTimerEventDefinition struct {
	ID           string              `xml:"id,attr,omitempty" json:"id"`
	TimeDate     []time.TimeDate     `xml:"timeDate,omitempty" json:"timeDate,omitempty"`
	TimeDuration []time.TimeDuration `xml:"timeDuration,omitempty" json:"timeDuration,omitempty"`
}

func NewTimerEventDefinition() TimerEventDefinitionRepository {
	return &TimerEventDefinition{}
}

/**
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (ted *TimerEventDefinition) SetID(typ string, suffix interface{}) {
	switch typ {
	case "ted":
		ted.ID = fmt.Sprintf("TimerEventDefinition_%v", suffix)
		break
	case "id":
		ted.ID = fmt.Sprintf("%s", suffix)
		break
	}
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
func (ted TimerEventDefinition) GetID() eventsbase.STR_PTR {
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
