package models

// Text ...
type Text struct {
	Text string `xml:",innerxml,omitempty"`
}

/* Content */

/** BPMN **/

// SetText ...
func (txt *Text) SetText(text string) {
	txt.Text = text
}
