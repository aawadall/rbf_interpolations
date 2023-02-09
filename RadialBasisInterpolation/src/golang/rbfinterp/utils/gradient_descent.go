package utils

import (
	"fmt"
)

// GradientDescent
func GradientDescent(A [][]float64, y []float64, epsilon float64, alpha float64) []float64 {
	// Initialize theta to 0 vector
	theta := make([]float64, len(A[0]))

	// Get objective function
	J := Objective(A, y, theta)

	// while J > epsilon
	for J > epsilon {
		fmt.Printf("J: %f\n", J)
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
