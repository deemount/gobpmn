package models

import "fmt"

// TimerEventDefinition ...
type TimerEventDefinition struct {
	ID       string   `xml:"id,attr,omitempty" json:"id"`
	TimeDate TimeDate `xml:"bpmn:timeDate,omitempty" json:"timeDate,omitempty"`
}

/* Attributes */

// SetID ...
func (ted *TimerEventDefinition) SetID(suffix string) {
	ted.ID = fmt.Sprintf("TimerEventDefinition_%s", suffix)
}
