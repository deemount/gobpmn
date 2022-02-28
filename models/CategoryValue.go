package models

// CategoryValue ...
type CategoryValue struct {
	ID    string `xml:"id,attr,omitempty" json:"id"`
	Value string `xml:"value,attr,omitempty" json:"value,omitempty"`
}
