package models

import "fmt"

// FlowNodeRefRepository ...
type FlowNodeRefRepository interface {
	SetID(typ string, suffix interface{})
}

// FlowNodeRef ...
type FlowNodeRef struct {
	ID string `xml:",innerxml,omitempty" json:"id"`
}

/**
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

/**
 * Default Getters
 */

/* Content */

// GetID ...
func (fnr FlowNodeRef) GetID() *string {
	return &fnr.ID
}
