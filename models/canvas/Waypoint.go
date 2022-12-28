package canvas

// NewWaypoint ...
func NewWaypoint() WaypointRepository {
	return &Waypoint{}
}

/*
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

/*
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
