package pool

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/loop"
	"github.com/deemount/gobpmn/models/marker"
)

// Collaboration ...
type Collaboration struct {
	compulsion.CoreID
	attributes.Attributes
	Participant PARTICIPANT_SLC      `xml:"bpmn:participant" json:"participant,omitempty"`
	MessageFlow []marker.MessageFlow `xml:"bpmn:messageFlow,omitempty" json:"messageFlow,omitempty"`
}

// TCollaboration ...
type TCollaboration struct {
	compulsion.CoreID
	attributes.TAttributes
	Participant TPARTICIPANT_SLC      `xml:"participant" json:"participant,omitempty"`
	MessageFlow []marker.TMessageFlow `xml:"messageFlow,omitempty" json:"messageFlow,omitempty"`
}

// FlowNodeRef ...
type FlowNodeRef struct {
	compulsion.CoreInnerID
}

// Lane ...
type Lane struct {
	compulsion.CoreID
	FlowNodeRef []FlowNodeRef `xml:"bpmn:flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

// TLane ...
type TLane struct {
	compulsion.CoreID
	FlowNodeRef []FlowNodeRef `xml:"flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

// LaneSet ...
type LaneSet struct {
	compulsion.CoreID
	Lane []Lane `xml:"bpmn:lane,omitempty" json:"lane,omitempty"`
}

// TLaneSet ...
type TLaneSet struct {
	compulsion.CoreID
	Lane []TLane `xml:"lane,omitempty" json:"lane,omitempty"`
}

// Participant ...
type Participant struct {
	compulsion.BaseAttributes
	ProcessRef              string                         `xml:"processRef,attr,omitempty" json:"processRef,omitempty" csv:"PROCESS_REF"`
	Documentation           []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty" csv:"DOCUMENTATION"`
	ParticipantMultiplicity []loop.ParticipantMultiplicity `xml:"bpmn:participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty" csv:"PARTICIPANT_MULTIPLICITY"`
}

// TParticipant ...
type TParticipant struct {
	compulsion.BaseAttributes
	ProcessRef              string                          `xml:"processRef,attr" json:"processRef,omitempty"`
	Documentation           []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ParticipantMultiplicity []loop.TParticipantMultiplicity `xml:"participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty"`
}
