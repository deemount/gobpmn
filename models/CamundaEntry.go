package models

// CamundaEntry ...
type CamundaEntry struct {
	Key   string `xml:"key,attr,omitempty"`
	Value string `xml:",innerxml,omitempty"`
}

/* Attributes */

/** Camunda **/

// SetKey ...
func (centry *CamundaEntry) SetKey(key string) {
	centry.Key = key
}

/* Content */

/** Camunda **/

// SetValue ...
func (centry *CamundaEntry) SetValue(value string) {
	centry.Value = value
}
