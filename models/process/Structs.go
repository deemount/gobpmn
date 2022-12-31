package process

import (
	"github.com/deemount/gobpmn/models/attributes"
	"github.com/deemount/gobpmn/models/compulsion"
	"github.com/deemount/gobpmn/models/data"
	"github.com/deemount/gobpmn/models/events"
	"github.com/deemount/gobpmn/models/gateways"
	"github.com/deemount/gobpmn/models/marker"
	"github.com/deemount/gobpmn/models/pool"
	"github.com/deemount/gobpmn/models/subprocesses"
	"github.com/deemount/gobpmn/models/tasks"
)

// Process ...
type Process struct {
	compulsion.BaseAttributes
	attributes.Attributes
	events.Events
	tasks.Tasks
	subprocesses.Subprocesses
	gateways.Gateways
	// Base
	IsExecutable bool `xml:"isExecutable,attr" json:"isExecutable,omitempty" csv:"IS_EXECUTABLE"`
	// Camunda
	CamundaVersionTag             string `xml:"camunda:versionTag,attr,omitempty" json:"versionTag,omitempty" csv:"VERSION_TAG"`
	CamundaJobPriority            int    `xml:"camunda:jobPriority,attr,omitempty" json:"jobPriority,omitempty" csv:"JOB_PRIORITY"`
	CamundaTaskPriority           int    `xml:"camunda:taskPriority,attr,omitempty" json:"taskPriority,omitempty" csv:"TASK_PRIORITY"`
	CamundaCandidateStarterGroups string `xml:"camunda:candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty" csv:"CANDIDATE_STARTER_GROUPS"`
	CamundaCandidateStarterUsers  string `xml:"camunda:candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty" csv:"CANDIDATE_STARTER_USERS"`
	CamundaHistoryTimeToLive      string `xml:"camunda:historyTimeToLive,attr,omitempty" json:"historyTimeToLive,omitempty" csv:"HISTORY_TIME_TO_LIVE"`
	// Pool
	LaneSet []pool.LaneSet `xml:"bpmn:laneSet,omitempty" json:"laneSet,omitempty" csv:"-"`
	// Marker
	Association  []marker.Association  `xml:"bpmn:association,omitempty" json:"association,omitempty"`
	SequenceFlow []marker.SequenceFlow `xml:"bpmn:sequenceFlow,omitempty" json:"sequenceFlow,omitempty" csv:"-"`
	Group        []marker.Group        `xml:"bpmn:group,omitempty" json:"group,omitempty" csv:"-"`
	// Data
	DataObject []data.DataObject `xml:"bpmn:dataObject,omitempty" json:"dataObject,omitempty"`
}

// TProcess ...
type TProcess struct {
	compulsion.BaseAttributes
	attributes.TAttributes
	events.TEvents
	tasks.TTasks
	subprocesses.TSubprocesses
	gateways.TGateways
	// Base
	IsExecutable bool `xml:"isExecutable,attr" json:"isExecutable,omitempty"`
	// Camunda
	VersionTag             string `xml:"versionTag,attr,omitempty" json:"versionTag,omitempty"`
	JobPriority            int    `xml:"jobPriority,attr,omitempty" json:"jobPriority,omitempty"`
	TaskPriority           int    `xml:"taskPriority,attr,omitempty" json:"taskPriority,omitempty"`
	CandidateStarterGroups string `xml:"candidateStarterGroups,attr,omitempty" json:"candidateStarterGroups,omitempty"`
	CandidateStarterUsers  string `xml:"candidateStarterUsers,attr,omitempty" json:"candidateStarterUsers,omitempty"`
	HistoryTimeToLive      string `xml:"historyTimeToLive,attr,omitempty" json:"historyTimeToLive,omitempty"`
	// Pool
	LaneSet []pool.LaneSet `xml:"laneSet,omitempty" json:"laneSet,omitempty"`
	// Marker
	Association  []marker.TAssociation  `xml:"association,omitempty" json:"association,omitempty"`
	SequenceFlow []marker.TSequenceFlow `xml:"sequenceFlow,omitempty" json:"sequenceFlow,omitempty"`
	Group        []marker.TGroup        `xml:"group,omitempty" json:"group,omitempty" csv:"-"`
	// Data
	DataObject []data.DataObject `xml:"dataObject,omitempty" json:"dataObject,omitempty"`
}
