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
func (ctl *CamundaTaskListener) SetClass(class string) {
	ctl.Class = class
}

// SetEvent ...
// can be: assignment, create, complete, delete, update, timeout
func (ctl *CamundaTaskListener) SetEvent(event string) {
	ctl.Event = event
}

// SetID ...
func (ctl *CamundaTaskListener) SetListenerID(listenerID string) {
	ctl.ListenerID = listenerID
}

/* Elements */

/** Camunda **/

// SetCamundaField ...
func (ctl *CamundaTaskListener) SetCamundaField(num int) {
	ctl.CamundaField = make([]CamundaField, num)
}
