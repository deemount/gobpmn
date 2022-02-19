package models

// Property ...
type Property struct {
	Name  string `xml:"name,attr,omitempty"`
	Value string `xml:"value,attr,omitempty"`
}

// SetName ...
func (prop *Property) SetName(name string) {
	prop.Name = name
}

// SetValue ...
func (prop *Property) SetValue(value string) {
	prop.Value = value
}
