package utils

import (
	"fmt"
	"math"

	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/types"
)

// type aliases
type Point = types.Point

// Square returns the square of coordinates
func Square(x Point) Point {
	// TODO - implement
	result := Point{
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
func Subtract(x, y Point) (Point, error) {

	// check that the points are the same dimensionality
	if x.Dimensionality != y.Dimensionality {
		fmt.Printf("x: %v", x.Dimensionality)
		fmt.Printf("y: %v", y.Dimensionality)
		return Point{}, fmt.Errorf("points are not the same dimensionality")
	}

	// TODO - implement
	result := Point{
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
func SquareRootSum(x Point) (float64, error) {
	// TODO - implement
	result := 0.0

	// loop through the coordinates
	for c := 0; c < x.Dimensionality; c++ {
		result += x.Coordinates[c] * x.Coordinates[c]
	}

	return math.Sqrt(result), nil
}
