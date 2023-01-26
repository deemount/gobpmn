package pool

import (
	"fmt"

	"github.com/deemount/gobpmn/models/bpmn/impl"
)

// NewFlowNodeRef ...
func NewFlowNodeRef() FlowNodeRefRepository {
	return &FlowNodeRef{}
}

/*
 * Default Setters
 */

/* Content */

// SetID ...
func (fnr *FlowNodeRef) SetID(typ string, suffix interface{}) {
	switch typ {
	case "id":
		fnr.ID = fmt.Sprintf("%s", suffix)
		break
	}
}

/*
 * Default Getters
 */

/* Content */

// GetID ...
func (fnr FlowNodeRef) GetID() impl.STR_PTR {
	return &fnr.ID
}
