package rbfinterp

import (
	"fmt"

	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/distances"
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/kernels"
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/types"
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/utils"
)

// type aliases
type RBFKernel = kernels.RBFKernel
type DistanceFunction = distances.DistanceFunction
type OptimizationFunction = utils.OptimizationFunction
type Point = types.Point

// RBFInterpolator is a struct that implements the Radial Basis Function Interpolation algorithm.
type RBFInterpolator struct {
	Configuration        map[string]interface{}
	Kernel               RBFKernel
	Distance             DistanceFunction
	OptimizationFunction OptimizationFunction
	KScale               float64
	SupportPoints        []Point
	SupportValues        []float64
	Weights              []float64
}

// Factory method for creating a new RBFInterpolator.
func NewRBFInterpolator(kernel RBFKernel, distance DistanceFunction, optimizationFunction OptimizationFunction, configuration map[string]interface{}) *RBFInterpolator {

	return &RBFInterpolator{
		Configuration:        configuration,
		Kernel:               kernel,
		Distance:             distance,
		OptimizationFunction: optimizationFunction,
		SupportPoints:        make([]Point, 0),
		SupportValues:        make([]float64, 0),
		Weights:              make([]float64, 0),
	}
}

// Load data into the RBFInterpolator.
func (r *RBFInterpolator) LoadData(x []Point, y []float64) error {
	// check that x and y are not nil
	if x == nil || y == nil {
		return fmt.Errorf("support points and their values must not be nil")
	}

	// check that x and y are not empty
	if len(x) == 0 || len(y) == 0 {
		return fmt.Errorf("support points and their values must not be empty")
	}
	// check that x and y are the same length
	if len(x) != len(y) {
		return fmt.Errorf("support points and their values must have same length")
	}

	// Load the data into the RBFInterpolator
	r.SupportPoints = make([]Point, len(x))
	r.SupportValues = make([]float64, len(y))
	for i := 0; i < len(x); i++ {
		r.SupportPoints[i] = x[i]
		r.SupportValues[i] = y[i]
	}

	return nil
}

// Train the RBFInterpolator.
func (r *RBFInterpolator) Train() error {
	fmt.Printf("Training RBF Interpolator")
	// check that the support points and values have been loaded
	if len(r.SupportPoints) == 0 || len(r.SupportValues) == 0 {
		return fmt.Errorf("support points and their values must be loaded before training")
	}

	// check that the support points and values are the same length
	if len(r.SupportPoints) != len(r.SupportValues) {
		return fmt.Errorf("support points and their values must have same length")
	}

	// build the similarity matrix of size n x n
	n := len(r.SupportPoints)
	//go func(i, j int) {
	//}(i, j)
	similarityMatrix := r.KMatrixParallel(n)
	fmt.Printf("Similarity Matrix Calculated")
	// TODO - calculate weights = inv(simTsim) * simT * y

	// find norm of similarity matrix
	r.KScale = utils.MatrixNorm(similarityMatrix)

	// scale K by norm
	similarityMatrix = utils.MatrixScale(similarityMatrix, 1.0/r.KScale)

	// initialize weights to random values
	r.Weights = make([]float64, n)
	for i := 0; i < n; i++ {
		r.Weights[i] = utils.RandomFloat64(-0.1, 0.1)
	}
	r.Weights = r.OptimizationFunction(similarityMatrix, r.SupportValues, r.Weights, r.Configuration)

	// inspect weights
	//fmt.Printf("\n\n\nWeights: %v, size = %d, support points size: %d", r.Weights, len(r.Weights), len(r.SupportPoints))
	return nil
}

func (r *RBFInterpolator) KMatrix(n int) [][]float64 {
	similarityMatrix := make([][]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Training RBF Interpolator: %d/%d\n", i, n)
		similarityMatrix[i] = make([]float64, n)
		for j := 0; j < n; j++ {

			similarityMatrix[i][j] = r.Kernel.Similarity(r.SupportPoints[i], r.SupportPoints[j])

		}
	}
	return similarityMatrix
}

// KMatrixParallel is a parallelized version of KMatrix.
func (r *RBFInterpolator) KMatrixParallel(n int) [][]float64 {
	similarityMatrix := make([][]float64, n)

	for i := 0; i < n; i++ {
		similarityMatrix[i] = make([]float64, n)
	}

	// TODO - parallelize this
	for i := 0; i < n; i++ {
		fmt.Printf("Training RBF Interpolator: %d/%d\n", i, n)
		for j := 0; j < n; j++ {
			go func(i, j int) {
				similarityMatrix[i][j] = r.Kernel.Similarity(r.SupportPoints[i], r.SupportPoints[j])
			}(i, j)
		}
	}

	return similarityMatrix
}

// Predict the value of a point.
func (r *RBFInterpolator) Predict(x Point) (float64, error) {
	fmt.Printf("Predicting value for point: %v", x)
	// check that the support points and values have been loaded
	if len(r.SupportPoints) == 0 || len(r.SupportValues) == 0 {
		fmt.Printf("support points and their values must be loaded before training")
		return 0, fmt.Errorf("support points and their values must be loaded before training")
	}

	// check that the support points and values are the same length
	if len(r.SupportPoints) != len(r.SupportValues) {
		fmt.Printf("support points and their values must have same length")
		return 0, fmt.Errorf("support points and their values must have same length")
	}

	// check that the weights have been calculated
	if len(r.Weights) == 0 {
		fmt.Printf("weights must be calculated before predicting")
		return 0, fmt.Errorf("weights must be calculated before predicting")
	}

	// check that the weights are the same length as the support points
	if len(r.Weights) != len(r.SupportPoints) {
		fmt.Printf("weights must be the same length as the support points")
		return 0, fmt.Errorf("weights must be the same length as the support points")
	}

	fmt.Printf("about to calculate similarity")
	// calculate similarity between x and each support point
	similarity := make([]float64, len(r.SupportPoints))
	for i := 0; i < len(r.SupportPoints); i++ {
		similarity[i] = r.Kernel.Similarity(x, r.SupportPoints[i])
	}

	// scale similarity by norm
	similarity = utils.VectorScale(similarity, 1.0/r.KScale)
	
	// print similarity
	//fmt.Printf("\n\nSimilarity: %v\n\n", similarity)

	fmt.Printf("about to calculate prediction")

	// calculate the prediction as the dot product of the weights and the similarity
	prediction := utils.DotProduct(r.Weights, similarity)

	// inspect prediction
	// fmt.Printf("\n\nPrediction: %f\n\n", prediction)
	return prediction, nil
}
