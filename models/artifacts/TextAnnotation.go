package artifacts

// TextAnnotationRepository ...
type TextAnnotationRepository interface {
	SetText()
	GetText() Text
}

// TextAnnotation ...
type TextAnnotation struct {
	Text []Text `xml:"bpmn:text,omitempty" json:"text,omitempty"`
}

// TTextAnnotation ...
type TTextAnnotation struct {
	Text []Text `xml:"text,omitempty" json:"text,omitempty"`
}

func NewTextAnnotation() TextAnnotationRepository {
	return &TextAnnotation{}
}

/**
 * Default Setters
 */

/*** Make Elements ***/

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
func (textannotation TextAnnotation) GetText() Text {
	return textannotation.Text[0]
}
