package marker

import "fmt"

// SignalRepository ...
type SignalRepository interface {
	MarkerBase
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

func NewSignal() SignalRepository {
	return &Signal{}
}

/*
 * Default Setters
 */

/* Attributes */

/** BPMN **/

// SetID ...
func (signal *Signal) SetID(typ string, suffix interface{}) {
	switch typ {
	case "signal":
		signal.ID = fmt.Sprintf("Signal_%v", suffix)
		break
	case "id":
		signal.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

// SetName ...
func (signal *Signal) SetName(suffix string) {
	signal.Name = fmt.Sprintf("Signal_%s", suffix)
}

/*
 * Default Getters
 */

/* Attributes */

/** BPMN **/

// GetID ...
func (signal Signal) GetID() STR_PTR {
	return &signal.ID
}

// GetName ...
func (signal Signal) GetName() STR_PTR {
	return &signal.Name
}
