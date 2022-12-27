package attributes

import "fmt"

// DocumentationRepository ...
type DocumentationRepository interface {
	SetData(data string)
	GetData() *string
}

// Documentation ...
type Documentation struct {
	Data string `xml:",innerxml,omitempty" json:"documentation,omitempty"`
}

// NewDocumenation ...
func NewDocumentation() DocumentationRepository {
	return &Documentation{}
}

/*
 * Default Setters
 */

/* Content */

// SetData ...
func (documentation *Documentation) SetData(data string) {
	documentation.Data = fmt.Sprintf("%s", data)
}

/*
 * Default Getters
 */

/* Content */

// GetData ...
func (documentation Documentation) GetData() *string {
	return &documentation.Data
}
