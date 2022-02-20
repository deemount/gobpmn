package models

import "fmt"

// SignalEventDefinition ...
type SignalEventDefinition struct {
	ID        string `xml:"id,attr,omitempty"`
	SignalRef string `xml:"signalRef,attr,omitempty"`
}

// SetID ...
func (sed *SignalEventDefinition) SetID(suffix string) {
	sed.ID = fmt.Sprintf("SignalEventDefinition_%s", suffix)
}

// SetSignalRef ...
func (sed *SignalEventDefinition) SetSignalRef(suffix string) {
	sed.SignalRef = fmt.Sprintf("Signal_%s", suffix)
}
