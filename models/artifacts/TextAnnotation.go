package artifacts

// NewTextAnnotation ...
func NewTextAnnotation() TextAnnotationRepository {
	return &TextAnnotation{}
}

/*
 * Default Setters
 */

/*** Make Elements ***/

/** BPMN **/

// SetText ...
func (textannotation *TextAnnotation) SetText() {
	textannotation.Text = make([]Text, 1)
}

/*
 * Default Getters
 */

/* Elements */

/** BPMN **/

// GetText ...
func (textannotation TextAnnotation) GetText() Text {
	return textannotation.Text[0]
}
