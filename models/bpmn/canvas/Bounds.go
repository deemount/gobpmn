package canvas

// NewBounds ...
func NewBounds() BoundsRepository {
	return &Bounds{}
}

/*
 * Default Setters
 */

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

/*
 * Default Getters
 */

// GetCoordinates ...
func (bnds Bounds) GetCoordinates() (*int, *int) {
	return &bnds.X, &bnds.Y
}

// GetX ...
func (bnds Bounds) GetX() *int {
	return &bnds.X
}

// GetY ...
func (bnds Bounds) GetY() *int {
	return &bnds.Y
}

// GetSize ...
func (bnds Bounds) GetSize() (*int, *int) {
	return &bnds.Width, &bnds.Height
}

// GetWidth ...
func (bnds Bounds) GetWidth() *int {
	return &bnds.Width
}

// GetHeight ...
func (bnds Bounds) GetHeight() *int {
	return &bnds.Height
}
