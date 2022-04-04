package models

// Label ...
type Label struct {
	Bounds []Bounds `xml:"dc:Bounds,omitempty" json:"-"`
}

/* Elements */

/** DC **/

// SetBounds ...
func (label *Label) SetBounds() {
	label.Bounds = make([]Bounds, 1)
}
