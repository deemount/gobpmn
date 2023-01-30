package camunda

import "github.com/deemount/gobpmn/models/bpmn/impl"

// NewCamundaTaskListener ...
func NewCamundaTaskListener() CamundaTaskListenerRepository {
	return &CamundaTaskListener{}
}

/*
 * Default Setters
 */

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

/*** Make Elements ***/

/** Camunda **/

// SetCamundaField ...
func (taskListener *CamundaTaskListener) SetCamundaField(num int) {
	taskListener.CamundaField = make([]CamundaField, num)
}

/*
 * Default Getters
 */

/* Attributes */

// GetClass ...
func (taskListener CamundaTaskListener) GetClass() impl.STR_PTR {
	return &taskListener.Class
}

// GetEvent ...
// can be: assignment, create, complete, delete, update, timeout
func (taskListener CamundaTaskListener) GetEvent() impl.STR_PTR {
	return &taskListener.Event
}

// GetID ...
func (taskListener CamundaTaskListener) GetListenerID() impl.STR_PTR {
	return &taskListener.ListenerID
}

/* Elements */

/** Camunda **/

// GetCamundaField ...
func (taskListener CamundaTaskListener) GetCamundaField(num int) *CamundaField {
	return &taskListener.CamundaField[num]
}
