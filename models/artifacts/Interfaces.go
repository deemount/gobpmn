package artifacts

// TextRepository ...
type TextRepository interface {
	SetContent(text string)
	GetContent() *string
}

// TextAnnotationRepository ...
type TextAnnotationRepository interface {
	SetText()
	GetText() Text
}
