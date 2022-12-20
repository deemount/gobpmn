package marker

import "fmt"

// CategoryValue ...
type CategoryValue struct {
	ID    string `xml:"id,attr,omitempty" json:"id"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID
func (categoryValue *CategoryValue) SetID(suffix string) {
	categoryValue.ID = fmt.Sprintf("_%s", suffix)
}

// SetValue
func (categoryValue *CategoryValue) SetValue(value string) {
	categoryValue.Value = value
}
