package canvas

import (
	"fmt"
	"log"
	"strconv"
)

// getDefaultCoordinates ...
func (p DelegateParameter) getDefaultCoordinates() (int, int) {
	var x, y int
	switch p.T {
	case "event":
		x, y = 179, 159
	case "pool":
		x, y = 600, 250
	case "startevent":
		x, y = 179, 159
	}
	return x, y
}

// getDefaultElementSize ...
func (p DelegateParameter) getDefaultElementSize() (int, int) {
	var width, height int
	switch p.T {
	case "activity":
		width, height = 100, 80
	case "event":
		width, height = 36, 36
	case "pool":
		width, height = 600, 250
	case "startevent":
		width, height = 36, 36

	}
	return width, height
}

// setBounds ...
func (p DelegateParameter) setBounds() {
	p.S.SetBounds()
	if p.B.X == 0 && p.B.Y == 0 {
		p.B.X, p.B.Y = p.getDefaultCoordinates()
	}
	p.S.GetBounds().SetCoordinates(p.B.X, p.B.Y)
	if p.B.Width == 0 && p.B.Height == 0 {
		p.B.Width, p.B.Height = p.getDefaultElementSize()
	}
	p.S.GetBounds().SetSize(p.B.Width, p.B.Height)
}

/*
 *
 */

// SetShape ...
// d *Shape, typ string, hash string, b Bounds
func SetShape(p DelegateParameter) {

	if p.S == nil && p.E == nil {
		log.Fatal("fatal: missing element pointer in canvas.SetShape()")
	}

	if p.T == "startevent" {
		// if startevent selected, then count up the _BPMNShape_StartEvent ID
		// NOTICE: Seen somewhere in the Camunda Modeller between Version 5 and 8.
		//         I really don't know, if it's making sense or for what the engine
		//         needs it later.
		// TODO: Find out, if it's needed to count up the Shape ID in this special case
		newHash, _ := strconv.Atoi(p.H)
		newHash += 1
		p.S.SetID(p.T, fmt.Sprint(newHash))

	} else {
		p.S.SetID(p.T, p.H)
	}

	p.S.SetElement(p.T, p.H)
	p.setBounds()
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
// e *Shape, typ string, isHorizontal bool, hash string, b Bounds
// e.g. below from collaborative process
// PoolCustomerSupport: canvas.Bounds{X: 150, Y: 80, Width: 800, Height: 160}
// PoolCustomer: canvas.Bounds{X: 150, Y: 360, Width: 800, Height: 160}
// CustomerSupportStartEvent: canvas.Bounds{X: 225, Y: 142, Width: 36, Height: 36}
// CustomerStartEvent: canvas.Bounds{X: 225, Y: 422, Width: 36, Height: 36}
// CustomerSupportEndEvent: canvas.Bounds{X: 822, Y: 142, Width: 36, Height: 36}
// CustomerEndEvent: canvas.Bounds{X: 822, Y: 422, Width: 36, Height: 36}
func SetPool(p DelegateParameter) {
	p.S.SetID(p.T, p.H)
	p.S.SetElement(p.T, p.H)
	p.S.SetIsHorizontal(p.I)
	p.S.SetBounds()
	// TODO:
	if &p.B.X == nil && &p.B.Y == nil && &p.B.Width == nil && &p.B.Height == nil {
		b := Bounds{X: 129, Y: 52, Width: 600, Height: 250}
		p.S.GetBounds().SetCoordinates(b.X, b.Y)
		p.S.GetBounds().SetSize(b.Width, b.Height)
	} else {
		p.S.GetBounds().SetCoordinates(p.B.X, p.B.Y)
		p.S.GetBounds().SetSize(p.B.Width, p.B.Height)
	}
}
