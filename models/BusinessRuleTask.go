package models

import "fmt"

type BusinessRuleTask struct {
	ID   string `xml:"id,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (brt *BusinessRuleTask) SetID(suffix string) {
	brt.ID = fmt.Sprintf("_%s", suffix)
}

// SetName ...
func (brt *BusinessRuleTask) SetName(name string) {
	brt.Name = name
}
