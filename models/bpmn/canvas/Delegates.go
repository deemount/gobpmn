package canvas

import (
	"fmt"
	"log"
	"strconv"
)

/*

	Some notices for the build of waypoints, bounds and other position utilities

	First, a rule set (axiomatically)

	There some ways to build a bpmn by a particular order.
	Let's try, to figurr out, which "build orders" are given

	1. Build Order:
		a) A Element has Flow in his method
			* will be drawn directly after the element
		b) A Element has no Flow in his method
			* will be drawn after elements were drawn

	TODO: Positioning of shapes and edges are in the making and not finished yet.

	A first handle to get in touch with the complexity is already implemented

*/

/*** Position ***/

// Center ...
type Center struct {
	x, y int
}

func center(bounds Bounds) Center {
	return Center{
		x: bounds.X + (bounds.Width / 2),
		y: bounds.Y + (bounds.Height / 2),
	}
}

// Big Delta calculates the difference between two elements
type Delta struct {
	x, y float64
}

func delta(a, b Point) Delta {
	return Delta{
		x: a.X - b.X,
		y: a.X - b.Y,
	}
}

/*** Defaults ***/

func (p DelegateParameter) defaultDistanceTop() int {
	return 50
}

func (p DelegateParameter) defaultDistanceLeft() int {
	return 50
}

// defaultEdgeLength ...
func (p DelegateParameter) defaultEdgeLength() int {
	return 55
}

// defaultCoordinates ...
func (p DelegateParameter) defaultCoordinates() (int, int) {
	var x, y int
	switch p.T {
	case "event":
		x, y = 179, 159
	case "pool":
		x, y = 600, 250
	case "startevent":
		x, y = 179, 159
	}
	log.Printf("canvas.delegates: set default coordinates x: %d, y: %d for %s", x, y, p.T)
	return x, y
}

// defaultElementSize ...
func (p DelegateParameter) defaultElementSize() (int, int) {

	var width, height int
	var typ string

	// option for the first waypoint, but collides with the second waypoint
	// TODO: change this condition. In this state it is too complex for usage
	if p.T == "flow" && p.ST != "" {
		typ = p.ST // the type, which is assigned the flow to
	} else {
		typ = p.T // the actually type
	}

	switch typ {
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

/*
 *
 */

// findCoordinatesByPreviousWaypoint ...
func (p *DelegateParameter) findCoordinatesByPreviousWaypoint() (int, int) {

	log.Println("canvas.delegates: find coordinates by previous waypoint no bounds set in shape")

	// To set the right coordinates for the given shape, I need to the height of the shape to calculate the point on the Y-axis
	// I assume, that the previous edge, which points to >>this<< shape, is a straight line from the last waypoint
	// So, p.WPPREV.X is the endpoint of the edge and connects directly to the shape, means it is the startpoint X of the shape

	log.Printf("canvas.delegates: find coordinates by previous waypoint got shape type %s", p.T)
	_, height := p.defaultElementSize()

	log.Printf("canvas.delegates: find coordinates by previous waypoint got previous waypoint %d, %d", p.WPPREV.X, p.WPPREV.Y)
	x := p.WPPREV.X
	y := p.WPPREV.Y - (height / 2)

	log.Printf("canvas.delegates: find coordinates by previous waypoint returns coordinates %d, %d", x, y)

	return x, y
}

// findCoordinatesByPreviousBounds ...
func (p *DelegateParameter) findCoordinatesByPreviousBounds() (int, int) {

	log.Println("canvas.delegates: find coordinates by previous bounds detects no bounds set in shape")

	// To set the right coordinates for the given shape, I need the height of the shape to calculate the point on the Y-axis
	// I assume, that the previous bounds, which is inside the shape (pool), has
	// So, p.WPPREV.X is the endpoint of the edge and connects directly to the shape, means it is the startpoint X of the shape

	log.Printf("canvas.delegates: find coordinates by previous bounds got shape type %s", p.T)
	_, height := p.defaultElementSize()

	log.Printf("canvas.delegates: find coordinates by previous bounds got previous waypoint %d, %d", p.BSPTR.X, p.BSPTR.Y)
	x := p.BSPTR.X + p.defaultDistanceLeft()
	y := (p.BSPTR.Y + (p.BSPTR.Height / 2)) - (height / 2)

	log.Printf("canvas.delegates: find coordinates by previous bounds returns coordinates %d, %d", x, y)

	return x, y
}

// setBounds sets the map for the shape,
// sets default or given coordinates for the shape
// and sets the elements default or given size of the shape
func (p *DelegateParameter) setBounds() {

	p.S.SetBounds()

	// if coordinates of x and y are zero, decide between two conditions:
	// * if previous waypoint is not nil, find the coordinates
	// * if
	// * else set default coordinates, when previous waypoint is nil
	if p.B.X == 0 && p.B.Y == 0 {

		if p.WPPREV != nil {
			p.B.X, p.B.Y = p.findCoordinatesByPreviousWaypoint()
		} else {
			if p.BSPTR != nil {
				p.B.X, p.B.Y = p.findCoordinatesByPreviousBounds()
			} else {
				p.B.X, p.B.Y = p.defaultCoordinates()
			}
		}

	}
	p.S.GetBounds().SetCoordinates(p.B.X, p.B.Y)

	// if width and height are zero, set default element size
	if p.B.Width == 0 && p.B.Height == 0 {
		p.B.Width, p.B.Height = p.defaultElementSize()
	}
	p.S.GetBounds().SetSize(p.B.Width, p.B.Height)

	log.Printf("canvas.delegates: set bounds for %s (%s)", p.S.ID, p.S.Element)

}

// setWaypoints ...
func (p *DelegateParameter) setWaypoints() {

	// Waypoint
	p.E.SetWaypoint()

	// waypoint sizes aren't given ...
	if len(p.W) == 0 {

		// calculate the first waypoint
		// e.g.
		// taking default values, which assigned above
		// startevent X1: 179
		// edge (from StartEvent) X2: unknown
		// startevent (Width) X3: 36
		// I want to know X2 and Y2 (where the edge should start)
		// X3 = X2 - X1 (get the size of startevent; base)
		// X2 - X1 = X3 | + X1 (formula conversion)
		// X2 = X3 + X1 (point where edge starts from)
		// X2 = 36 + 179
		// X2 = 215
		// startevent Y1: 159
		// edge (from StartEvent) Y2: unknown
		// startevent (Height) Y3: 36
		// Y2 = startevent Y1 + startevent (Height) / 2
		// Y2 = 159 + (36 / 2)
		// Y2 = 177
		// Notice: I need to know, which shape is before the start of a waypoint
		//		   BSPTR is the pointer to the element before
		width, height := p.defaultElementSize()   // X3, Y3
		x1 := width + p.BSPTR.X                   // X2
		y1 := (height / 2) + p.BSPTR.Y            // Y2
		p.E.GetWaypoint(0).SetCoordinates(x1, y1) // e.g. []canvas.Waypoint{{X: 215, Y: 177}, ...}

		// calculate the second waypoint
		// an example of how to calculate is given above. For the second waypoint, other values needs to be recognised.
		//
		// TODO: Do I need to know, which shape is after the start of a waypoint?
		//
		// Answer  I: If I follow the build order with default sizes, the second waypoint describes the bounds of the next element
		//            The next element is then in dependence of the flow, showing to.
		//			  For this option I need to know the >>default<< length of a flow and X of the first waypoint.
		//			  Also, I need to know, if the element is at the same Y-Position like the element before.
		//			  Like for X value, I need to use the build order I explained at the top of the page
		//			  I assume that the following element is at the same Y-position from where the previous elements flow is starting from.
		//			  e.g.
		//			  startevent waypoint 2 X1: unknown
		//			  default edge length X2: 55
		//			  X of first waypoint X3: 215
		//			  X1 = X2 + X3
		//			  X1 = 55 + 215
		//			  X1 = 270
		//
		// Answer II: If I don't follow the build order, which means all flows >>must<< set after all elements.
		//			  Therefore I need the bounds of the following element (which isn't actually set in the delegate parameter)

		// code for answer I
		x2 := p.defaultEdgeLength() + x1
		y2 := y1
		p.E.GetWaypoint(1).SetCoordinates(x2, y2) // e.g. []canvas.Waypoint{..., {X: 270, Y: 177}}

	} else {

		p.E.GetWaypoint(0).SetCoordinates(p.W[0].X, p.W[0].Y)
		p.E.GetWaypoint(1).SetCoordinates(p.W[1].X, p.W[1].Y)

	}

	log.Printf("canvas.delegates: set waypoints for %s", p.E.ID)

}

/*
 *
 */

// SetShape ...
func SetShape(p DelegateParameter) {

	if p.S == nil && p.E == nil {
		log.Fatal("fatal: missing element pointer in canvas.SetShape()")
	}

	// detect startevent
	//
	if p.T == "startevent" {

		// TODO: Find out, if it's needed to count up the Shape ID in this special case (Camunda related)
		// if startevent selected, then count up the _BPMNShape_StartEvent ID
		// NOTICE: Seen somewhere in the Camunda Modeller between Version 5 and 8.
		//         I really don't know, if it's making sense or for what the engine
		//         needs it later.
		newHash, _ := strconv.Atoi(p.H)
		newHash += 1
		p.S.SetID(p.T, fmt.Sprint(newHash))

	} else {
		p.S.SetID(p.T, p.H)
	}

	p.S.SetElement(p.T, p.H)
	p.setBounds()

	log.Printf("canvas.delegates: set shape for %s (%s)", p.S.ID, p.S.Element)

}

// SetEdge ...
func SetEdge(p DelegateParameter) {
	p.E.SetID(p.T, p.H)
	p.E.SetElement(p.T, p.H)
	p.setWaypoints()
	log.Printf("canvas.delegates: set edge for %s (%s)", p.E.ID, p.E.Element)
}

// SetLabel ...
func SetLabel(p DelegateParameter) {
	p.E.SetLabel()
	p.E.GetLabel().SetBounds()
	p.E.GetLabel().GetBounds().SetCoordinates(p.B.X, p.B.Y)
	p.E.GetLabel().GetBounds().SetSize(p.B.Width, p.B.Height)
}

// SetPool ...
// default coords and size are: x="129" y="52" width="600" height="250"
// when start event is also default at x="179" y="159" width="36" height="36"
// Notice:
// width of pool >>must<< grow with length of the process, which is defined with
// an endevent (terminate and so on) at his X,Y-points and a default distance right
// to the element.
// calculating X and Y for a startevent inside the first pool:
// e.g.
// startevent X = pool X + default distance left
// pool X = startevent X - default distance left
// default distance left = startevent X - pool X
//
// e.g.
// startevent Y = ((pool (height) / 2) + pool Y) - (startevent (height) / 2)
// (startevent (height) / 2) = ((pool (height) / 2) + pool Y) - startevent Y
//
// e *Shape, typ string, isHorizontal bool, hash string, b Bounds
// e.g. below from collaborative process
// TODO: recalculate the real X and Y by given defaults from pool at
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

	defaultCoordinates := Bounds{X: 129, Y: 52}
	defaultSize := Bounds{Width: 600, Height: 250}

	// TODO: calculate X,Y for many pools

	// If a pool has a bounds pointer, it can't be the first pool
	// The bounds pointer in a pool set shows the next pool the orientation
	if p.BSPTR != nil {

		if p.B.X == 0 && p.B.Y == 0 {
			if p.BSPTR.X == defaultCoordinates.X && p.BSPTR.Y == defaultCoordinates.Y {
				x := defaultCoordinates.X
				y := defaultCoordinates.Y + defaultSize.Height + p.defaultDistanceTop()
				p.S.GetBounds().SetCoordinates(x, y)
			}
		}

	} else {

		if p.B.X == 0 && p.B.Y == 0 {
			p.S.GetBounds().SetCoordinates(defaultCoordinates.X, defaultCoordinates.Y)
		} else {
			p.S.GetBounds().SetCoordinates(p.B.X, p.B.Y)
		}

	}

	// TODO: calculate Width,Height for many pools
	if p.B.Width == 0 && p.B.Height == 0 {
		p.S.GetBounds().SetSize(defaultSize.Width, defaultSize.Height)
	} else {
		p.S.GetBounds().SetSize(p.B.Width, p.B.Height)
	}

	log.Printf("canvas.delegates: set bounds for %s (%s)", p.S.ID, p.S.Element)

}
