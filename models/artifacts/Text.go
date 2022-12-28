package artifacts

import "fmt"

// TextRepository ...
type TextRepository interface {
	SetContent(text string)
	GetContent() *string
}

// Text ...
type Text struct {
	Text string `xml:",innerxml,omitempty" json:"text,omitempty"`
}

// NewText ...
func NewText() TextRepository {
	return &Text{}
}

/**
 * Default Setters
 */

/* Content */

/** BPMN **/

// SetText ...
func (txt *Text) SetContent(text string) {
	txt.Text = fmt.Sprintf("%s", text)
}

/**
 * Default Getters
 */

/* Content */

/** BPMN **/

// GetText ...
func (txt Text) GetContent() *string {
	return &txt.Text
}
