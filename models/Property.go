package models

// Property ...
type Property struct {
	Name  string `xml:"name,attr,omitempty" json:"name,omitempty"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

/* Attributes */

// SetName ...
func (prop *Property) SetName(name string) {
	prop.Name = name
}

// SetValue ...
func (prop *Property) SetValue(value string) {
	prop.Value = value
}
