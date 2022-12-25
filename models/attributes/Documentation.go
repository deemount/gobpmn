package attributes

import "fmt"

// DocumentationRepository ...
type DocumentationRepository interface {
	SetData(data string)
	GetData() string
}

// Documentation ...
type Documentation struct {
	Data string `xml:",innerxml,omitempty" json:"elementDocumentation,omitempty"`
}

func NewDocumentation() DocumentationRepository {
	return &Documentation{}
}

/* Content */

// SetData ...
func (documentation *Documentation) SetData(data string) {
	documentation.Data = fmt.Sprintf("%s", data)
}

/* Content */

// GetData ...
func (documentation Documentation) GetData() string {
	return documentation.Data
}
