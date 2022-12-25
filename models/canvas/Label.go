package canvas

// Label ...
type LabelRepository interface {
	SetBounds()
	GetBounds() *Bounds
}

// Label ...
type Label struct {
	Bounds []Bounds `xml:"dc:Bounds,omitempty" json:"-"`
}

func NewLabel() LabelRepository {
	return &Label{}
}

/**
 * Default Setters
 */

/*** Make Elements ***/

/** DC **/

// SetBounds ...
func (label *Label) SetBounds() {
	label.Bounds = make([]Bounds, 1)
}

/**
 * Default Getters
 */

/* Elements */

/** DC **/

// GetBounds ...
func (label Label) GetBounds() *Bounds {
	return &label.Bounds[0]
}
