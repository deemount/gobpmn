package models

import "fmt"

// Bounds ...
type Bounds struct {
	X      string `xml:"x,attr,omitempty"`
	Y      string `xml:"y,attr,omitempty"`
	Width  string `xml:"width,attr,omitempty"`
	Height string `xml:"height,attr,omitempty"`
}

// SetCoordinates ...
func (bnds *Bounds) SetCoordinates(x, y int64) {
	bnds.X = fmt.Sprintf("%d", x)
	bnds.Y = fmt.Sprintf("%d", y)
}

// SetX ...
func (bnds *Bounds) SetX(x int64) {
	bnds.X = fmt.Sprintf("%d", x)
}

// SetY ...
func (bnds *Bounds) SetY(y int64) {
	bnds.Y = fmt.Sprintf("%d", y)
}

// SetSize ...
func (bnds *Bounds) SetSize(width, height int64) {
	bnds.Width = fmt.Sprintf("%d", width)
	bnds.Height = fmt.Sprintf("%d", height)
}

// SetWidth ...
func (bnds *Bounds) SetWidth(width int64) {
	bnds.Width = fmt.Sprintf("%d", width)
}

// SetHeight ...
func (bnds *Bounds) SetHeight(height int64) {
	bnds.Height = fmt.Sprintf("%d", height)
}
