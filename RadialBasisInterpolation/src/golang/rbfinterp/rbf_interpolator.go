package rbfinterp

import (
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/distances"
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/kernels"
)

// type aliases
type RBFKernel = kernels.RBFKernel
type DistanceFunction = distances.DistanceFunction

// RBFInterpolator is a struct that implements the Radial Basis Function Interpolation algorithm.
type RBFInterpolator struct {
	Kernel   RBFKernel
	Distance DistanceFunction
}

// Factory method for creating a new RBFInterpolator.
func NewRBFInterpolator(kernel RBFKernel, distance DistanceFunction) *RBFInterpolator {
	// TODO - implement
	return &RBFInterpolator{
		Kernel:   kernel,
		Distance: distance,
	}
}
