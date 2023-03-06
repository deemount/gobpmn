package canvas_test

import (
	"log"
	"math"
	"testing"
)

func DistanceTest(t *testing.T) {

	type Point struct {
		X float64
		Y float64
	}

	type Line struct {
		P1 Point
		P2 Point
	}

	line := Line{
		Point{float64(215), float64(237)},
		Point{float64(300), float64(237)},
	}

	distance := math.Sqrt(math.Pow(line.P2.X-line.P1.X, 2) + math.Pow(line.P2.Y-line.P1.Y, 2))

	log.Printf("main: euclidean distance is %.2f\n", distance)
}
