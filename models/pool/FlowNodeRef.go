package pool

import "fmt"

// FlowNodeRefRepository ...
type FlowNodeRefRepository interface {
	PoolID
}

// FlowNodeRef ...
type FlowNodeRef struct {
	ID string `xml:",innerxml,omitempty" json:"id"`
}

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
func (fnr FlowNodeRef) GetID() STR_PTR {
	return &fnr.ID
}
