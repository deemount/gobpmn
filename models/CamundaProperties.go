package models

// CamundaProperties ...
type CamundaProperties struct {
	Name  string `xml:"name,attr,omitempty"`
	Value string `xml:"value,attr,omitempty"`
}

// SetName ...
func (cproperties *CamundaProperties) SetName(name string) {
	cproperties.Name = name
}

// SetValue ...
func (cproperties *CamundaProperties) SetValue(value string) {
	cproperties.Value = value
}
