package models

// CamundaTaskListener ...
type CamundaTaskListener struct {
	Class        string         `xml:"class,attr,omitempty" json:"class,omitempty"`
	Event        string         `xml:"event,attr,omitempty" json:"event,omitempty"`
	ListenerID   string         `xml:"id,attr,omitempty" json:"listenerId,omitempty"`
	CamundaField []CamundaField `xml:"camunda:field,omitempty" json:"field,omitempty"`
}

/* Attributes */

// SetClass ...
func (taskListener *CamundaTaskListener) SetClass(class string) {
	taskListener.Class = class
}

// SetEvent ...
// can be: assignment, create, complete, delete, update, timeout
func (taskListener *CamundaTaskListener) SetEvent(event string) {
	taskListener.Event = event
}

// SetID ...
func (taskListener *CamundaTaskListener) SetListenerID(listenerID string) {
	taskListener.ListenerID = listenerID
}

/* Elements */

/** Camunda **/

// SetCamundaField ...
func (taskListener *CamundaTaskListener) SetCamundaField(num int) {
	taskListener.CamundaField = make([]CamundaField, num)
}
