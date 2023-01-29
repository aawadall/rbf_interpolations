package distances

import (
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp"
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/utils"
)

// type aliases 
type DistanceFunction func(x, y rbfinterp.Point) (float64, error)
type Point = rbfinterp.Point

// function aliases
var (
	Square = utils.Square
	SquareRootSum = utils.SquareRootSum
	Subtract  = utils.Subtract
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
