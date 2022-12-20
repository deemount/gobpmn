package marker

import "fmt"

// SignalRepository ...
type SignalRepository interface {
	SetID(suffix string)
	SetName(name string)
	GetID() string
	GetName() string
}

// Signal ...
type Signal struct {
	ID   string `xml:"id,attr,omitempty" json:"id"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}

// TSignal ...
type TSignal struct {
	ID   string `xml:"id,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID ...
func (signal *Signal) SetID(suffix string) {
	signal.ID = fmt.Sprintf("Signal_%s", suffix)
}

// SetName ...
func (signal *Signal) SetName(suffix string) {
	signal.Name = fmt.Sprintf("Signal_%s", suffix)
}

// SetID ...
func (signal Signal) GetID() string {
	return signal.ID
}

// SetName ...
func (signal Signal) GetName() string {
	return signal.Name
}
