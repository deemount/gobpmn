package models

import (
	"fmt"
)

// Waypoint ...
type Waypoint struct {
	X string `xml:"x,attr"`
	Y string `xml:"y,attr"`
}

// SetX ...
func (wp *Waypoint) SetX(x int64) {
	wp.X = fmt.Sprintf("%d", x)
}

// SetY ...
func (wp *Waypoint) SetY(y int64) {
	wp.Y = fmt.Sprintf("%d", y)
}

// SetCoordinates ...
func (wp *Waypoint) SetCoordinates(x, y int64) {
	wp.X = fmt.Sprintf("%d", x)
	wp.Y = fmt.Sprintf("%d", y)
}
