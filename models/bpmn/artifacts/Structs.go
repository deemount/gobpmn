package artifacts

// Text ...
type Text struct {
	Text string `xml:",innerxml,omitempty" json:"text,omitempty"`
}

// TextAnnotation ...
type TextAnnotation struct {
	Text []Text `xml:"bpmn:text,omitempty" json:"text,omitempty"`
}

// TTextAnnotation ...
type TTextAnnotation struct {
	Text []Text `xml:"text,omitempty" json:"text,omitempty"`
}
