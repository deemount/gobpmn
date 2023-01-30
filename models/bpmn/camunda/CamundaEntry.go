package camunda

import "github.com/deemount/gobpmn/models/bpmn/impl"

// NewCamundaEntry ...
func NewCamundaEntry() CamundaEntryRepository {
	return &CamundaEntry{}
}

/*
 * Default Setters
 */

/* Attributes */

// SetKey ...
func (centry *CamundaEntry) SetKey(key string) {
	centry.Key = key
}

/* Content */

// SetValue ...
func (centry *CamundaEntry) SetValue(value string) {
	centry.Value = value
}

/*
 * Default Getters
 */

/* Attributes */

// GetKey ...
func (centry CamundaEntry) GetKey() impl.STR_PTR {
	return &centry.Key
}

/* Content */

// GetValue ...
func (centry CamundaEntry) GetValue() impl.STR_PTR {
	return &centry.Value
}
