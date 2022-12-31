package repository

import (
	"time"

	"github.com/deemount/gobpmn/models/core"
	"github.com/deemount/gobpmn/spec/process_instance"
)

// ProcessInfo ...
type ProcessInfo struct {
	BpmnProcessId string            // The ID as defined in the BPMN file
	Version       int32             // A version of the process, default=1, incremented, when another process with the same ID is loaded
	ProcessKey    int64             // The engines key for this given process with version
	ModelType     string            // The type of the model. Can be human-driven or system-driven
	Def           core.TDefinitions // parsed file content
	checksumBytes [16]byte          // internal checksum to identify different versions
}

// ProcessInstanceInfo ...
type ProcessInstanceInfo struct {
	processInfo     *ProcessInfo
	instanceKey     int64
	variableContext map[string]interface{}
	createdAt       time.Time
	state           process_instance.State
}
