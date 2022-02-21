package models

import "fmt"

// EventBasedGateway ...
type EventBasedGateway struct {
	ID                string              `xml:"id,attr"`
	Name              string              `xml:"name,attr,omitempty"`
	ExtensionElements []ExtensionElements `xml:"bpmn:extensionElements,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (eventbasedgate *EventBasedGateway) SetID(suffix string) {
	eventbasedgate.ID = fmt.Sprintf("Gateway_%s", suffix)
}

// SetName ...
func (eventbasedgate *EventBasedGateway) SetName(name string) {
	eventbasedgate.Name = name
}
