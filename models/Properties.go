package models

// Properties ...
type Properties struct {
	Property []Property `xml:"bpmn:property,omitempty"`
}

// SetProperty ...
func (properties *Properties) SetProperty(num int) {
	properties.Property = make([]Property, num)
}
