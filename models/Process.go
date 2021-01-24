package models

// Process ...
type Process struct {
	ID                string `xml:"id,attr"`
	IsExecutable      string `xml:"isExecutable,attr"`
	CamundaVersionTag string `xml:"camunda:versionTag,attr,omitempty"`
}

// SetID ...
func (proc *Process) SetID() {
	proc.ID = "Process_"
}

// SetIsExecutable ...
func (proc *Process) SetIsExecutable() {
	proc.IsExecutable = "true"
}

// SetCamundaVersionTag ...
func (proc *Process) SetCamundaVersionTag() {
	proc.CamundaVersionTag = "0.1"
}
