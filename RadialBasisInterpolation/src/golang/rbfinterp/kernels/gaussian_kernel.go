package kernels

import (
	"math"

	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/distances"
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/types"
)

// type aliases
type Point = types.Point
type DistanceFunction = distances.DistanceFunction

type GaussianKernel struct {
	DistanceFunction DistanceFunction
	Parameters       map[string]interface{}
}

// GaussianKernel is a struct that implements the Gaussian Radial Basis Function.
func NewGaussianKernel(distance DistanceFunction, params map[string]interface{}) *GaussianKernel {
	// TODO - implement
	return &GaussianKernel{
		DistanceFunction: distance,
		Parameters:       params,
	}
}

// similarity returns the similarity between two points.
func (gk *GaussianKernel) similarity(x, y Point) float64 {
	// TODO - implement
	distance, err := gk.DistanceFunction(x, y)

	if err != nil {
		panic(err)
	}

	sigma := gk.Parameters["sigma"].(float64)

	return math.Exp(-math.Pow(distance, 2.0) / (2.0 * math.Pow(sigma, 2.0)))

}
