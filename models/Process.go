package models

// Process ...
type Process struct {
	ID                     string                 `xml:"id,attr"`
	Name                   string                 `xml:"name,attr,omitempty"`
	IsExecutable           string                 `xml:"isExecutable,attr"`
	CamundaVersionTag      string                 `xml:"camunda:versionTag,attr,omitempty"`
	StartEvent             StartEvent             `xml:"bpmn:startEvent"`
	IntermediateThrowEvent IntermediateThrowEvent `xml:"bpmn:intermediateThrowEvent"`
	SequenceFlow           SequenceFlow           `xml:"bpmn:sequenceFlow"`
}

// SetID ...
func (proc *Process) SetID(suffix string) {
	proc.ID = "Process_" + suffix
}

// SetName ...
func (proc *Process) SetName(name string) {
	proc.Name = name
}

// SetIsExecutable ...
func (proc *Process) SetIsExecutable() {
	proc.IsExecutable = "true"
}

// SetCamundaVersionTag ...
func (proc *Process) SetCamundaVersionTag(tag string) {
	proc.CamundaVersionTag = tag
}
