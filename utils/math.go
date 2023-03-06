package utils

import "math"

func FixError(n float64) float64 {
	return math.Round(n*100000000000) / 100000000000
}

// ToFixed formats a number using fixed-point notation.
func ToFixed(n float64, precision int) float64 {
	scale := math.Pow(10, float64(precision))
	return math.Round(n*scale) / scale
}
