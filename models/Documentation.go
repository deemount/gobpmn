package models

import "fmt"

// Documentation ...
type Documentation struct {
	ElementDocumentation string `xml:",innerxml,omitempty" json:"elementDeocumentation,omitempty"`
}

/* Content */

// SetElementDocumentation ...
func (documentation *Documentation) SetElementDocumentation(elementDocu string) {
	documentation.ElementDocumentation = fmt.Sprintf("%s", elementDocu)
}
