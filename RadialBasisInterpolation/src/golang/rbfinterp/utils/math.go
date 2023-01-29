package utils

import (
	"fmt"
	"math"

	"github.com/aawadall/rbf_interpolations/golang/rbfinterp"
)

// Square returns the square of coordinates
func Square(x rbfinterp.Point) rbfinterp.Point {
	// TODO - implement
	result := rbfinterp.Point{
		Dimensionality: x.Dimensionality,
		Coordinates:    make([]float64, x.Dimensionality),
	}

	// loop through the coordinates
	for c := 0; c < x.Dimensionality; c++ {
		result.Coordinates[c] = x.Coordinates[c] * x.Coordinates[c]
	}

	return result
}

// Subtract returns the difference between two points
func Subtract(x, y rbfinterp.Point) (rbfinterp.Point, error) {

	// check that the points are the same dimensionality
	if x.Dimensionality != y.Dimensionality {
		return rbfinterp.Point{}, fmt.Errorf("points are not the same dimensionality")
	}

	// TODO - implement
	result := rbfinterp.Point{
		Dimensionality: x.Dimensionality,
		Coordinates:    make([]float64, x.Dimensionality),
	}

	// loop through the coordinates
	for c := 0; c < x.Dimensionality; c++ {
		result.Coordinates[c] = x.Coordinates[c] - y.Coordinates[c]
	}

	return result, nil
}

// SquareRootSum returns the square root of the sum of the squares of the coordinates
func SquareRootSum(x rbfinterp.Point) (float64, error) {
	// TODO - implement
	result := 0.0

	// loop through the coordinates
	for c := 0; c < x.Dimensionality; c++ {
		result += x.Coordinates[c] * x.Coordinates[c]
	}

	return math.Sqrt(result), nil
}
