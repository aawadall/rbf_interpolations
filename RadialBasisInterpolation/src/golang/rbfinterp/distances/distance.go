package distances

import (
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/types"
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/utils"
)

// type aliases
type Point = types.Point
type DistanceFunction func(x, y Point) (float64, error)

// function aliases
var (
	Square        = utils.Square
	SquareRootSum = utils.SquareRootSum
	Subtract      = utils.Subtract
)

// EuclideanDistance
func EuclideanDistance(x, y Point) (float64, error) {
	// return sqrt(Sum(x^2 - y^2))
	difference, err := Subtract(x, y)
	if err != nil {
		return 0.0, err
	}
	squared_diff := Square(difference)

	return SquareRootSum(squared_diff)

}
