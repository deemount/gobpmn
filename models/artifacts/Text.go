package artifacts

import "fmt"

// TextRepository ...
type TextRepository interface {
	SetText(text string)
}

// Text ...
type Text struct {
	Text string `xml:",innerxml,omitempty" json:"text,omitempty"`
}

func NewText() TextRepository {
	return &Text{}
}

/**
 * Default Setters
 */

/* Content */

/** BPMN **/

// SetText ...
func (txt *Text) SetText(text string) {
	txt.Text = fmt.Sprintf("%s", text)
}

/**
 * Default Getters
 */

/* Content */

/** BPMN **/

// GetText ...
func (txt *Text) GetText() *string {
	return &txt.Text
}
