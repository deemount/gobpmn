package camunda

// CamundaValueRepository ...
type CamundaValueRepository interface {
	CamundaBaseValue
}

// CamundaValue ...
type CamundaValue struct {
	Value string `xml:",innerxml,omitempty" json:"value,omitempty"`
}

// NewCamundaValue ...
func NewCamundaValue() CamundaValueRepository {
	return &CamundaValue{}
}

/*
 * Default Setters
 */

/* Content */

// SetValue ...
func (value *CamundaValue) SetValue(val string) {
	value.Value = val
}

/*
 * Default Getters
 */

/* Content */

// GetValue ...
func (value CamundaValue) GetValue() STR_PTR {
	return &value.Value
}
