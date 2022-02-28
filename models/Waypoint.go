package models

// Waypoint ...
type Waypoint struct {
	X int `xml:"x,attr" json:"-"`
	Y int `xml:"y,attr" json:"-"`
}

// SetX ...
func (wp *Waypoint) SetX(x int) {
	wp.X = x
}

// SetY ...
func (wp *Waypoint) SetY(y int) {
	wp.Y = y
}

// SetCoordinates ...
func (wp *Waypoint) SetCoordinates(x, y int) {
	wp.X = x
	wp.Y = y
}
