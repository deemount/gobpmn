package models

import "fmt"

// Category ...
type Category struct {
	ID string `xml:"id,attr,omitempty"`
}

/* Attributes */

/** BPMN **/

// SetID
func (category *Category) SetID(suffix string) {
	category.ID = fmt.Sprintf("_%s", suffix)
}
