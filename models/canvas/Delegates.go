package canvas

// SetShape ...
// d *Shape, typ string, hash string, b Bounds
func SetShape(p DelegateParameter) {
	p.S.SetID(p.T, p.H)
	p.S.SetElement(p.T, p.H)
	p.S.SetBounds()
	p.S.GetBounds().SetCoordinates(p.B.X, p.B.Y)
	p.S.GetBounds().SetSize(p.B.Width, p.B.Height)
}

// SetEdge ...
// d *Edge, typ string, hash string, w []Waypoint
func SetEdge(p DelegateParameter) {
	p.E.SetID(p.T, p.H)
	p.E.SetElement(p.T, p.H)
	p.E.SetWaypoint()
	p.E.GetWaypoint(0).SetCoordinates(p.W[0].X, p.W[0].Y)
	p.E.GetWaypoint(1).SetCoordinates(p.W[1].X, p.W[1].Y)
}

// SetLabel ...
// d *Edge, b Bounds
func SetLabel(p DelegateParameter) {
	p.E.SetLabel()
	p.E.GetLabel().SetBounds()
	p.E.GetLabel().GetBounds().SetCoordinates(p.B.X, p.B.Y)
	p.E.GetLabel().GetBounds().SetSize(p.B.Width, p.B.Height)
}

// SetPool ...
// default coords and size are: x="129" y="52" width="600" height="250"
// when start event is also default at x="179" y="159" width="36" height="36"
func SetPool(e *Shape, typ string, isHorizontal bool, hash string, b Bounds) {
	e.SetID(typ, hash)
	e.SetElement(typ, hash)
	e.SetIsHorizontal(isHorizontal)
	e.SetBounds()
	if &b.X == nil && &b.Y == nil && &b.Width == nil && &b.Height == nil {
		b := Bounds{X: 129, Y: 52, Width: 600, Height: 250}
		e.GetBounds().SetCoordinates(b.X, b.Y)
		e.GetBounds().SetSize(b.Width, b.Height)
	} else {
		e.GetBounds().SetCoordinates(b.X, b.Y)
		e.GetBounds().SetSize(b.Width, b.Height)
	}
}
