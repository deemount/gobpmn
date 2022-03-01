package models

import "fmt"

// Documentation ...
type Documentation struct {
	Data string `xml:",innerxml,omitempty" json:"elementDeocumentation,omitempty"`
}

/* Content */

// SetElementDocumentation ...
func (documentation *Documentation) SetData(data string) {
	documentation.Data = fmt.Sprintf("%s", data)
}
