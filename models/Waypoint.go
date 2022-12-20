package models

// WaypointRepository ...
type WaypointRepository interface {
	SetX(x int)
	SetY(y int)
	SetCoordinates(x, y int)

	GetX() *int
	GetY() *int
	GetCoordinates() (*int, *int)
}

// Waypoint ...
type Waypoint struct {
	X int `xml:"x,attr" json:"-"`
	Y int `xml:"y,attr" json:"-"`
}

/**
 * Default Setters
 */

/* Attributes */

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

/**
 * Default Getters
 */

/* Attributes */

// GetX ...
func (wp Waypoint) GetX() *int {
	return &wp.X
}

// GetY ...
func (wp Waypoint) GetY() *int {
	return &wp.Y
}

// GetCoordinates ...
func (wp Waypoint) GetCoordinates() (*int, *int) {
	return &wp.X, &wp.Y
}
