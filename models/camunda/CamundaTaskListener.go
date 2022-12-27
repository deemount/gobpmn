package camunda

// CamundaTaskListener ...
type CamundaTaskListenerRepository interface {
	CamundaBaseEvent
	CamundaBaseClass
	SetListenerID(listenerID string)
	GetListenerID() *string
	SetCamundaField(num int)
	GetCamundaField(num int) *CamundaField
}

// CamundaTaskListener ...
type CamundaTaskListener struct {
	Class        string         `xml:"class,attr,omitempty" json:"class,omitempty"`
	Event        string         `xml:"event,attr,omitempty" json:"event,omitempty"`
	ListenerID   string         `xml:"id,attr,omitempty" json:"listenerId,omitempty"`
	CamundaField []CamundaField `xml:"camunda:field,omitempty" json:"field,omitempty"`
}

// TCamundaTaskListener ...
type TCamundaTaskListener struct {
	Class      string         `xml:"class,attr,omitempty" json:"class,omitempty"`
	Event      string         `xml:"event,attr,omitempty" json:"event,omitempty"`
	ListenerID string         `xml:"id,attr,omitempty" json:"listenerId,omitempty"`
	Field      []CamundaField `xml:"field,omitempty" json:"field,omitempty"`
}

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
func (taskListener CamundaTaskListener) GetClass() STR_PTR {
	return &taskListener.Class
}

// GetEvent ...
// can be: assignment, create, complete, delete, update, timeout
func (taskListener CamundaTaskListener) GetEvent() STR_PTR {
	return &taskListener.Event
}

// GetID ...
func (taskListener CamundaTaskListener) GetListenerID() *string {
	return &taskListener.ListenerID
}

/* Elements */

/** Camunda **/

// GetCamundaField ...
func (taskListener CamundaTaskListener) GetCamundaField(num int) *CamundaField {
	return &taskListener.CamundaField[num]
}
