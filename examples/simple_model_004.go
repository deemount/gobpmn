package examples

import (
	"github.com/deemount/gobpmn/models"
	"github.com/deemount/gobpmn/utils"
)

// SimpleModel004Repository ...
type SimpleModel004Repository interface {
	Create()
	SetElements()
	SetDefinitionsAttributes()
	SetCollaboration()
	SetProcess()
	SetStartEvent()
	SetStartEventState()
	SetTask()
	SetTaskState()
	SetEndEvent()
	SetFromStartEventSequenceFlow()
	SetFromStartEventStateSequenceFlow()
	SetFromTaskSequenceFlow()
	SetFromTaskStateSequenceFlow()
	SetDiagram()
}

// SimpleModel004 ...
type SimpleModel004 struct {
	def                 *models.Definitions
	isExecutable        bool
	CollaborationHash   string
	ParticipantHash     string
	ProcessHash         string
	StartEventCounter   int64
	FromStartEvent      string
	FromStartEventState string
	TaskHash            string
	FromTask            string
	FromTaskState       string
	EndEventHash        string
}

// NewSimpleModel004 ...
func NewSimpleModel004(def *models.Definitions) SimpleModel004 {
	return SimpleModel004{
		def:               def,
		isExecutable:      true,
		CollaborationHash: utils.GenerateHash(),
		ParticipantHash:   utils.GenerateHash(),
		ProcessHash:       utils.GenerateHash(),
		StartEventCounter: 1,
		FromStartEvent:    utils.GenerateHash(),
		TaskHash:          utils.GenerateHash(),
		EndEventHash:      utils.GenerateHash(),
	}
}

// Create ...
func (simpleModel004 *SimpleModel004) Create() {}

// SetElements ...
func (simpleModel004 *SimpleModel004) SetElements() {}

// SetDefinitionsAttributes ...
func (simpleModel004 *SimpleModel004) SetDefinitionsAttributes() {
	simpleModel004.def.SetDefaultAttributes()
}

// SetProcess ...
func (simpleModel004 *SimpleModel004) SetProcess() {}

// SetStartEvent ...
func (simpleModel004 *SimpleModel004) SetStartEvent() {}

// SetStartEventState ...
func (simpleModel004 *SimpleModel004) SetStartEventState() {}

// SetTask ...
func (simpleModel004 *SimpleModel004) SetTask() {}

// SetTaskState ...
func (simpleModel004 *SimpleModel004) SetTaskState() {}

// SetEndEvent ...
func (simpleModel004 *SimpleModel004) SetEndEvent() {}

// SetFromStartEventSequenceFlow ...
func (simpleModel004 *SimpleModel004) SetFromStartEventSequenceFlow() {}

// SetFromStartEventStateSequenceFlow ...
func (simpleModel004 *SimpleModel004) SetFromStartEventStateSequenceFlow() {}

// SetFromTaskSequenceFlow ...
func (simpleModel004 *SimpleModel004) SetFromTaskSequenceFlow() {}

// SetFromTaskStateSequenceFlow ...
func (simpleModel004 *SimpleModel004) SetFromTaskStateSequenceFlow() {}

// SetDiagram ...
func (simpleModel004 *SimpleModel004) SetDiagram() {}
