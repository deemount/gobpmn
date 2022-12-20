package events

import (
	"fmt"

	"github.com/deemount/gobpmn/models/attributes"
)

// TimerEventDefinitionRepository ...
type TimerEventDefinitionRepository interface {
	SetID(typ string, suffix interface{})
	SetTimeDate()
	SetTimeDuration()

	GetID() *string

	GetTimeDate() *attributes.TimeDate
	GetTimeDuration() *attributes.TimeDuration
}

// TimerEventDefinition ...
type TimerEventDefinition struct {
	ID           string                    `xml:"id,attr,omitempty" json:"id"`
	TimeDate     []attributes.TimeDate     `xml:"bpmn:timeDate,omitempty" json:"timeDate,omitempty"`
	TimeDuration []attributes.TimeDuration `xml:"bpmn:timeDuration,omitempty" json:"timeDuration,omitempty"`
}

// TTimerEventDefinition ...
type TTimerEventDefinition struct {
	ID           string                    `xml:"id,attr,omitempty" json:"id"`
	TimeDate     []attributes.TimeDate     `xml:"timeDate,omitempty" json:"timeDate,omitempty"`
	TimeDuration []attributes.TimeDuration `xml:"timeDuration,omitempty" json:"timeDuration,omitempty"`
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
	ted.TimeDate = make([]attributes.TimeDate, 1)
}

// SetTimeDuration ...
func (ted *TimerEventDefinition) SetTimeDuration() {
	ted.TimeDuration = make([]attributes.TimeDuration, 1)
}

/**
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (ted TimerEventDefinition) GetID() *string {
	return &ted.ID
}

/* Elements */

/** BPMN **/

// GetTimeDuration ...
func (ted TimerEventDefinition) GetTimeDate() *attributes.TimeDate {
	return &ted.TimeDate[0]
}

// GetTimeDuration ...
func (ted TimerEventDefinition) GetTimeDuration() *attributes.TimeDuration {
	return &ted.TimeDuration[0]
}
