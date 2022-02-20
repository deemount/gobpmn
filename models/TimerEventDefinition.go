package models

import "fmt"

// TimerEventDefinition ...
type TimerEventDefinition struct {
	ID       string   `xml:"id,attr,omitempty"`
	TimeDate TimeDate `xml:"bpmn:timeDate,omitempty"`
}

// SetID ...
func (ted *TimerEventDefinition) SetID(suffix string) {
	ted.ID = fmt.Sprintf("TimerEventDefinition_%s", suffix)
}
