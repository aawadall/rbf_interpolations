package kernels

// unit tests

import (
	"math"
	"testing"

	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/distances"
)

// TestNewGaussianKernel
func TestNewGaussianKernel(t *testing.T) {
	// happy path
	params := map[string]interface{}{
		"sigma": 1.0,
	}
	euclideanDistance := distances.EuclideanDistance
	gk := NewGaussianKernel(euclideanDistance, params)

	if gk.DistanceFunction == nil {
		t.Errorf("NewGaussianKernel() => DistanceFunction is nil")
	}

	if gk.Parameters == nil {
		t.Errorf("NewGaussianKernel() => Parameters is nil")
	}

	if gk.Parameters["sigma"] != 1.0 {
		t.Errorf("NewGaussianKernel() => Parameters[\"sigma\"] != 1.0")
	}

}

// TestGaussianKernelSimilarity
func TestGaussianKernelSimilarity(t *testing.T) {
	// happy path
	params := map[string]interface{}{
		"sigma": 0.5,
	}
	distance := distances.EuclideanDistance
	gk := NewGaussianKernel(distance, params)

	x := Point{Dimensionality: 2, Coordinates: []float64{0.0, 0.0}}
	y := Point{Dimensionality: 2, Coordinates: []float64{1.0, 1.0}}

	similarity := gk.similarity(x, y)

	difference := make([]float64, 2)
	for i := 0; i < 2; i++ {
		difference[i] = x.Coordinates[i] - y.Coordinates[i]
	}

	squareDiff := make([]float64, 2)
	for i := 0; i < 2; i++ {
		squareDiff[i] = math.Pow(difference[i], 2.0)
	}

	sum := 0.0
	for i := 0; i < 2; i++ {
		sum += squareDiff[i]
	}

	expected := math.Exp(-sum / (2.0 * math.Pow(params["sigma"].(float64), 2.0)))

	// approximate equality
	if math.Abs(similarity-expected) > 0.0000000000000001 {
		t.Errorf("GaussianKernel.similarity() => similarity != expected; %f != %f", similarity, expected)
	}
}
