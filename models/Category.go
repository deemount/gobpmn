package models

import "fmt"

// Category ...
type Category struct {
	ID string `xml:"id,attr,omitempty" json:"id"`
}

/* Attributes */

/** BPMN **/

// SetID
func (category *Category) SetID(suffix string) {
	category.ID = fmt.Sprintf("_%s", suffix)
}
