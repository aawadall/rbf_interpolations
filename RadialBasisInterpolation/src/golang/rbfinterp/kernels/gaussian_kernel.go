package kernels

import (
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/types"
)

// type aliases
type Point = types.Point
type GaussianKernel struct {
}

// GaussianKernel is a struct that implements the Gaussian Radial Basis Function.
func NewGaussianKernel(params map[string]interface{}) *GaussianKernel {
	// TODO - implement
	return &GaussianKernel{}
}

// similarity returns the similarity between two points.
func (gk *GaussianKernel) similarity(x, y Point) float64 {
	// TODO - implement
	return 0.0
}
