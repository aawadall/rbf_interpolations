package scratchpad

import "math"

// GradientDescent
func GradientDescent(A [][]float64, y []float64, epsilon float64) []float64 {
	// Initialize theta to 0 vector
	theta := make([]float64, len(A[0]))

	// set alpha 
	alpha := 0.01

	// Get objective function
	J := Objective(A, y, theta)
	
	// while J > epsilon
	for J > epsilon {
		// Calculate the gradient of J with respect to theta
		gradient := Gradient(A, y, theta)

		// Update theta
		for i := 0; i < len(theta); i++ {
			theta[i] -= alpha * gradient[i]
		}

		// Update J
		J = Objective(A, y, theta)
	}

	return theta
}

// Objective
func Objective(A [][]float64, y []float64, theta []float64) float64 {
	// Given dense matrix A, vector y, and vector theta
	// Calculate y_hat as A * theta
	y_hat := MatVecDot(A, theta)

	// Calculate errVector as y - y_hat
	errVector := make([]float64, len(y))
	for i := 0; i < len(y); i++ {
		errVector[i] = y[i] - y_hat[i]
	}

	// Calculate objective as L2 norm of error
	J := L2Norm(errVector)

	return J
}

func L2Norm(vector []float64) float64 {
	J := 0.0
	for i := 0; i < len(vector); i++ {
		J += math.Pow(vector[i], 2.0)
	}
	return J
}

func MatVecDot(A [][]float64, theta []float64) []float64 {
	y_hat := make([]float64, len(A))
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			y_hat[i] += A[i][j] * theta[j]
		}
	}
	return y_hat
}


// Gradient
func Gradient(A [][]float64, y []float64, theta []float64) []float64 {
	// Given dense matrix A, vector y, and vector theta
	// Calculate y_hat as A * theta
	y_hat := MatVecDot(A, theta)
