package models

// FlowNodeRef ...
type FlowNodeRef struct {
	ID string `xml:",innerxml,omitempty"`
}

/* Content */

// SetID ...
func (fnr *FlowNodeRef) SetID(id string) {
	fnr.ID = id
}
