package kernels

import (
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp"
)

type RBFKernel interface {
	// similarity returns the similarity between two points.
	similarity(x, y rbfinterp.Point) float64
}

// Main Factory method for creating a new RBFKernel.
func NewRBFKernel(kernelType string, parameters map[string]interface{}) RBFKernel {
	// TODO - implement a switch statement to return the appropriate kernel given the kernel type and kernel parameters

	switch kernelType {
	case "gaussian":
		return NewGaussianKernel(parameters)
	}
	return nil
}

