package kernels

type RBFKernel interface {
	// Similarity returns the Similarity between two points.
	Similarity(x, y Point) float64
}

// Main Factory method for creating a new RBFKernel.
func NewRBFKernel(kernelType string, distance DistanceFunction, parameters map[string]interface{}) RBFKernel {
	// TODO - implement a switch statement to return the appropriate kernel given the kernel type and kernel parameters

	switch kernelType {
	case "gaussian":
		return NewGaussianKernel(distance, parameters)
	}
	return nil
}
