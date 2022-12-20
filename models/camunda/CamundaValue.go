package camunda

// CamundaValue ...
type CamundaValue struct {
	Value string `xml:",innerxml,omitempty" json:"value,omitempty"`
}

/* Content */

// SetValue ...
func (value *CamundaValue) SetValue(val string) {
	value.Value = val
}
