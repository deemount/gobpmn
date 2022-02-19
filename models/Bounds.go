package models

// Bounds ...
type Bounds struct {
	X      int `xml:"x,attr,omitempty"`
	Y      int `xml:"y,attr,omitempty"`
	Width  int `xml:"width,attr,omitempty"`
	Height int `xml:"height,attr,omitempty"`
}

// SetCoordinates ...
func (bnds *Bounds) SetCoordinates(x, y int) {
	bnds.X = x
	bnds.Y = y
}

// SetX ...
func (bnds *Bounds) SetX(x int) {
	bnds.X = x
}

// SetY ...
func (bnds *Bounds) SetY(y int) {
	bnds.Y = y
}

// SetSize ...
func (bnds *Bounds) SetSize(width, height int) {
	bnds.Width = width
	bnds.Height = height
}

// SetWidth ...
func (bnds *Bounds) SetWidth(width int) {
	bnds.Width = width
}

// SetHeight ...
func (bnds *Bounds) SetHeight(height int) {
	bnds.Height = height
}
