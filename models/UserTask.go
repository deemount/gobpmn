package models

type UserTask struct {
	ID                  string     `xml:"id,attr"`
	Name                string     `xml:"name,attr,omitempty"`
	CamundaFormKey      string     `xml:"camunda:formKey,attr,omitempty"`
	CamundaCandidUsers  string     `xml:"camunda:candidateUsers,attr,omitempty"`
	CamundaCandidGroups string     `xml:"camunda:candidateGroups,attr,omitempty"`
	CamundaPriority     int        `xml:"camunda:priority,attr,omitempty"`
	Incoming            []Incoming `xml:"bpmn:incoming,omitempty"`
	Outgoing            []Outgoing `xml:"bpmn:outgoing,omitempty"`
}

// SetID ...
func (ut *UserTask) SetID(suffix string) {
	ut.ID = "Activity_" + suffix
}

// SetName ...
func (ut *UserTask) SetName(name string) {
	ut.Name = name
}

/* Camunda */

// SetCamundaFormKey ...
func (ut *UserTask) SetCamundaFormKey(formKey string) {
	ut.CamundaFormKey = formKey
}

// SetCamundaCandidUsers ...
func (ut *UserTask) SetCamundaCandidUsers(users string) {
	ut.CamundaCandidUsers = users
}

// SetCamundaCandidGroups ...
func (ut *UserTask) SetCamundaCandidGroups(groups string) {
	ut.CamundaCandidGroups = groups
}

// SetCamundaPriority ...
func (ut *UserTask) SetCamundaPriority(priority int) {
	ut.CamundaPriority = priority
}
