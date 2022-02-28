package models

// FlowNodeRef ...
type FlowNodeRef struct {
	ID string `xml:",innerxml,omitempty" json:"id"`
}

/* Content */

// SetID ...
func (fnr *FlowNodeRef) SetID(id string) {
	fnr.ID = id
}
