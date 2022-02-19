package models

// Task ...
type Task struct {
	ID       string     `xml:"id,attr"`
	Name     string     `xml:"name,attr,omitempty"`
	Incoming []Incoming `xml:"bpmn:incoming,omitempty"`
	Outgoing []Outgoing `xml:"bpmn:outgoing,omitempty"`
}

// SetID ...
func (tk *Task) SetID(suffix string) {
	tk.ID = "Activity_" + suffix
}

// SetName ...
func (tk *Task) SetName(name string) {
	tk.Name = name
}

// SetIncoming ...
func (tk *Task) SetIncoming(num int) {
	tk.Incoming = make([]Incoming, num)
}

// SetOutgoing ...
func (tk *Task) SetOutgoing(num int) {
	tk.Outgoing = make([]Outgoing, num)
}
