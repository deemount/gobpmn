package models

// CamundaTaskListener ...
type CamundaTaskListener struct {
	Class        string         `xml:"class,attr,omitempty"`
	Event        string         `xml:"event,attr,omitempty"`
	ID           string         `xml:"id,attr,omitempty"`
	CamundaField []CamundaField `xml:"camunda:field,omitempty"`
}

/* Attributes */

// SetClass ...

// SetEvent ...

// SetID ...

/* Elements */

/** Camunda **/

// SetCamundaField
