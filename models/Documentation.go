package models

import "fmt"

// Documentation ...
type Documentation struct {
	ElementDocumentation string `xml:",innerxml,omitempty"`
}

/* Content */

// SetElementDocumentation ...
func (docu *Documentation) SetElementDocumentation(elementDocu string) {
	docu.ElementDocumentation = fmt.Sprintf("%s", elementDocu)
}
