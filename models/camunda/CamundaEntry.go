package camunda

// CamundaEntryRepository ...
type CamundaEntryRepository interface {
	CamundaBaseValue
	SetKey(key string)
	GetKey() *string
}

// CamundaEntry ...
type CamundaEntry struct {
	Key   string `xml:"key,attr,omitempty" json:"key"`
	Value string `xml:",innerxml,omitempty" json:"value,omitempty"`
}

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
func (centry CamundaEntry) GetKey() *string {
	return &centry.Key
}

/* Content */

// GetValue ...
func (centry CamundaEntry) GetValue() STR_PTR {
	return &centry.Value
}
