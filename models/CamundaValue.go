package models

// CamundaValue ...
type CamundaValue struct {
	Value string `xml:",innerxml,omitempty"`
}

/* Content */

// SetValue ...
func (cval *CamundaValue) SetValue(value string) {
	cval.Value = value
}
