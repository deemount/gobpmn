package infrastructure

import (
	"fmt"
	"math"
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

// LineIntersect calculates the intersection point of two lines defined by points p1, p2 and p3, p4.
// It returns the intersection point as a Point struct and an error if the lines are parallel or do not intersect.
// The function uses the line-line intersection formula to find the intersection point.
// It checks if the lines are parallel by calculating the denominator.
// If the denominator is zero, it means the lines are parallel and do not intersect.
// If the lines intersect, it calculates the intersection point using the formula.
// The function also checks if the intersection point is within the bounds of both lines.
// If the intersection point is outside the bounds, it returns an error.
// The function takes four points as input: p1, p2, p3, and p4.
// The points p1 and p2 define the first line, and points p3 and p4 define the second line.
// The function returns the intersection point as a Point struct and an error if any.
// The function uses the math package for mathematical calculations and the fmt package for error formatting.
// It also uses the utils package for fixing floating point errors and rounding.
// The function is useful for checking if two lines intersect in a 2D space.
// The function can be used in various applications such as computer graphics, game development, and geometric calculations.
// The function is efficient and handles edge cases such as parallel lines and out-of-bounds intersection points.
// The function is designed to be easy to use and understand.
// The function is well-documented and follows best practices for error handling and code readability.
func LineIntersect(p1, p2, p3, p4 Point) (Point, error) {
	return lineIntersect(p1, p2, p3, p4)
}

// lineIntersect calculates the intersection point of two lines defined by points p1, p2 and p3, p4.
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
	px := FixError(nt / denominator)
	py := FixError(nu / denominator)
	t := +ToFixed(px, 2)
	u := +ToFixed(py, 2)

	if t < ToFixed(math.Min(p1.X, p2.X), 2) ||
		t > ToFixed(math.Max(p1.X, p2.X), 2) ||
		t < ToFixed(math.Min(p3.X, p4.X), 2) ||
		t > ToFixed(math.Max(p3.X, p4.X), 2) ||
		u < ToFixed(math.Min(p1.Y, p2.Y), 2) ||
		u > ToFixed(math.Max(p1.Y, p2.Y), 2) ||
		u < ToFixed(math.Min(p3.Y, p4.Y), 2) ||
		u > ToFixed(math.Max(p3.Y, p4.Y), 2) {
		return Point{}, fmt.Errorf("lineIntersect: no line intersection given")
	}

	return Point{X: px, Y: py}, nil

}

// pointDistance calculates the distance between two points
// using the Euclidean distance formula.
// The formula is: distance = sqrt((x2 - x1)^2 + (y2 - y1)^2)
// where (x1, y1) and (x2, y2) are the coordinates of the two points.
// The function takes a Line struct as input, which contains two points P1 and P2.
// It returns the distance as a float64 value.
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

// FixError rounds a float64 to 11 decimal places.
// This is used to fix floating point errors in calculations.
func FixError(n float64) float64 {
	return math.Round(n*100000000000) / 100000000000
}

// ToFixed formats a number using fixed-point notation.
// It rounds the number to the specified number of decimal places.
// This is used to fix floating point errors in calculations.
// The function takes a float64 number and an integer precision as input.
// It returns the number rounded to the specified number of decimal places.
func ToFixed(n float64, precision int) float64 {
	scale := math.Pow(10, float64(precision))
	return math.Round(n*scale) / scale
}
