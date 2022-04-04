package models

import "fmt"

// TimerEventDefinition ...
type TimerEventDefinition struct {
	ID           string         `xml:"id,attr,omitempty" json:"id"`
	TimeDate     []TimeDate     `xml:"bpmn:timeDate,omitempty" json:"timeDate,omitempty"`
	TimeDuration []TimeDuration `xml:"bpmn:timeDuration,omitempty" json:"timeDuration,omitempty"`
}

/* Attributes */

// SetID ...
func (ted *TimerEventDefinition) SetID(suffix string) {
	ted.ID = fmt.Sprintf("TimerEventDefinition_%s", suffix)
}

// SetTimeDuration ...
func (ted *TimerEventDefinition) SetTimeDate() {
	ted.TimeDate = make([]TimeDate, 1)
}

// SetTimeDuration ...
func (ted *TimerEventDefinition) SetTimeDuration() {
	ted.TimeDuration = make([]TimeDuration, 1)
}
