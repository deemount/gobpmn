package pool

import (
<<<<<<< HEAD
	"github.com/deemount/gobpmn/models/bpmn/impl"
=======
	"github.com/deemount/gobpmn/models/bpmn/attributes"
	"github.com/deemount/gobpmn/models/bpmn/flow"
	"github.com/deemount/gobpmn/models/bpmn/impl"
	"github.com/deemount/gobpmn/models/bpmn/loop"
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
)

// pool
type DelegateParameter struct {
<<<<<<< HEAD
	T string   // typ
	N string   // name
	H []string // hash
=======
	PPT *Participant // element pointer
	T   string       // typ
	PR  string       // typProcessRef
	N   string       // name
	H   []string     // hash
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
}

/*
 * Elementary
 */

<<<<<<< HEAD
=======
// Collaboration ...
type Collaboration struct {
	impl.CoreID
	attributes.Attributes
	Participant PARTICIPANT_SLC    `xml:"bpmn:participant" json:"participant,omitempty"`
	MessageFlow []flow.MessageFlow `xml:"bpmn:messageFlow,omitempty" json:"messageFlow,omitempty"`
}

// TCollaboration ...
type TCollaboration struct {
	impl.CoreID
	attributes.TAttributes
	Participant TPARTICIPANT_SLC    `xml:"participant" json:"participant,omitempty"`
	MessageFlow []flow.TMessageFlow `xml:"messageFlow,omitempty" json:"messageFlow,omitempty"`
}

>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
// FlowNodeRef ...
type FlowNodeRef struct {
	impl.CoreInnerID
}

// Lane ...
type Lane struct {
	impl.CoreID
	FlowNodeRef []FlowNodeRef `xml:"bpmn:flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

// TLane ...
type TLane struct {
	impl.CoreID
	FlowNodeRef []FlowNodeRef `xml:"flowNodeRef,omitempty" json:"flowNodeRef,omitempty"`
}

// LaneSet ...
type LaneSet struct {
	impl.CoreID
	Lane []Lane `xml:"bpmn:lane,omitempty" json:"lane,omitempty"`
}

// TLaneSet ...
type TLaneSet struct {
	impl.CoreID
	Lane []TLane `xml:"lane,omitempty" json:"lane,omitempty"`
}
<<<<<<< HEAD
=======

// Participant ...
type Participant struct {
	impl.BaseAttributes
	ProcessRef              string                         `xml:"processRef,attr,omitempty" json:"processRef,omitempty" csv:"PROCESS_REF"`
	Documentation           []attributes.Documentation     `xml:"bpmn:documentation,omitempty" json:"documentation,omitempty" csv:"DOCUMENTATION"`
	ParticipantMultiplicity []loop.ParticipantMultiplicity `xml:"bpmn:participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty" csv:"PARTICIPANT_MULTIPLICITY"`
}

// TParticipant ...
type TParticipant struct {
	impl.BaseAttributes
	ProcessRef              string                          `xml:"processRef,attr" json:"processRef,omitempty"`
	Documentation           []attributes.Documentation      `xml:"documentation,omitempty" json:"documentation,omitempty"`
	ParticipantMultiplicity []loop.TParticipantMultiplicity `xml:"participantMultiplicity,omitempty" json:"participantMultiplicity,omitempty"`
}
>>>>>>> 645a2223b8faa7522808edfd0136a0ac06a3e7f4
