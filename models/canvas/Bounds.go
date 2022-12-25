package canvas

// BoundsRepository ...
type BoundsRepository interface {
	SetCoordinates(x, y int)
	SetX(x int)
	SetY(y int)
	SetSize(width, height int)
	SetWidth(width int)
	SetHeight(height int)

	GetCoordinates() (*int, *int)
	GetX() *int
	GetY() *int
	GetSize() (*int, *int)
	GetWidth() *int
	GetHeight() *int
}

// Bounds ...
type Bounds struct {
	X      int `xml:"x,attr,omitempty"`
	Y      int `xml:"y,attr,omitempty"`
	Width  int `xml:"width,attr,omitempty"`
	Height int `xml:"height,attr,omitempty"`
}

func NewBounds() BoundsRepository {
	return &Bounds{}
}

/**
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

/**
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
