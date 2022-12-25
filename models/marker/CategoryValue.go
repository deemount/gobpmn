package marker

import "fmt"

// CategoryValue ...
type CategoryValueRepository interface {
	MarkerID

	SetValue(value string)
	GetValue() *string
}

// CategoryValue ...
type CategoryValue struct {
	ID    string `xml:"id,attr,omitempty" json:"id"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID
func (categoryValue *CategoryValue) SetID(typ string, suffix interface{}) {
	switch typ {
	case "":
		categoryValue.ID = fmt.Sprintf("_%s", suffix)
		break
	case "id":
		categoryValue.ID = fmt.Sprintf("_%s", suffix)
		break
	}
}

// SetValue
func (categoryValue *CategoryValue) SetValue(value string) {
	categoryValue.Value = value
}

/* Attributes */

/** BPMN **/

// GetID
func (categoryValue CategoryValue) GetID() *string {
	return &categoryValue.ID
}

// GetValue
func (categoryValue CategoryValue) GetValue() *string {
	return &categoryValue.Value
}
