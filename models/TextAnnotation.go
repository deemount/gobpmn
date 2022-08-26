package models

// TextAnnotation ...
type TextAnnotation struct {
	Text []Text `xml:"bpmn:text,omitempty" json:"text,omitempty"`
}

// TTextAnnotation ...
type TTextAnnotation struct {
	Text []Text `xml:"bpmn:text,omitempty" json:"text,omitempty"`
}

/**
 * Default Setters
 */

/* Elements */

/** BPMN **/

// SetText ...
func (textannotation *TextAnnotation) SetText() {
	textannotation.Text = make([]Text, 1)
}

/**
 * Default Getters
 */

/* Elements */

/** BPMN **/

// GetText ...
func (textannotation TextAnnotation) GetText() *Text {
	return &textannotation.Text[0]
}
