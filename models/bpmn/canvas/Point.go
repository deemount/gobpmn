package canvas

import (
	"fmt"
	"math"

	"github.com/deemount/gobpmn/utils"
)

// Point ...
type Point struct {
	X float64
	Y float64
}

// Line ...
type Line struct {
	P1 Point
	P2 Point
}

func lineIntersect(p1, p2, p3, p4 Point) (Point, error) {

	if math.Max(p1.X, p2.X) < math.Min(p3.X, p4.X) ||
		math.Min(p1.X, p2.X) > math.Max(p3.X, p4.X) ||
		math.Max(p1.Y, p2.Y) < math.Min(p3.Y, p4.Y) ||
		math.Min(p1.Y, p2.Y) > math.Max(p3.Y, p4.Y) {
		return Point{}, fmt.Errorf("lineIntersect: point relations not given")
	}

	// calculate relative location of the intersection t and u
	nt := (p1.X*p2.Y-p1.Y*p2.X)*(p3.X-p4.X) - (p1.X-p2.X)*(p3.X*p4.Y-p3.Y*p4.X)
	nu := (p1.X*p2.Y-p1.Y*p2.X)*(p3.Y-p4.Y) - (p1.Y-p2.Y)*(p3.X*p4.Y-p3.Y*p4.X)

	// Denomination
	denominator := (p1.X-p2.X)*(p3.Y-p4.Y) - (p1.Y-p2.Y)*(p3.X-p4.X)
	if math.IsNaN(denominator) || denominator == 0 {
		return Point{}, fmt.Errorf("lineIntersect: lines are parallel")
	}

	// fxing Inf and NaN
	px := utils.FixError(nt / denominator)
	py := utils.FixError(nu / denominator)
	t := +utils.ToFixed(px, 2)
	u := +utils.ToFixed(py, 2)

	if t < +utils.ToFixed(math.Min(p1.X, p2.X), 2) ||
		t > +utils.ToFixed(math.Max(p1.X, p2.X), 2) ||
		t < +utils.ToFixed(math.Min(p3.X, p4.X), 2) ||
		t > +utils.ToFixed(math.Max(p3.X, p4.X), 2) ||
		u < +utils.ToFixed(math.Min(p1.Y, p2.Y), 2) ||
		u > +utils.ToFixed(math.Max(p1.Y, p2.Y), 2) ||
		u < +utils.ToFixed(math.Min(p3.Y, p4.Y), 2) ||
		u > +utils.ToFixed(math.Max(p3.Y, p4.Y), 2) {
		return Point{}, fmt.Errorf("lineIntersect: no line intersection given")
	}

	return Point{X: px, Y: py}, nil

}

func pointDistance(l1 Line) float64 {
	return math.Sqrt(math.Pow(l1.P2.X-l1.P1.X, 2) + math.Pow(l1.P2.Y-l1.P1.Y, 2))
}

// pointsOnLine returns true if the point r is on the line between p and q
func pointsOnLine(line Line, r Point) bool {

	accuracy := float64(5)

	value := (line.P2.X-line.P1.X)*(r.Y-line.P1.Y) - (line.P2.Y-line.P1.Y)*(r.X-line.P1.X)
	dist := pointDistance(line)

	// @see http://stackoverflow.com/a/907491/412190
	return math.Abs(value/dist) <= accuracy

}
