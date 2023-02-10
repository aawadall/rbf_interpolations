package utils

import (
	"fmt"
)

type OptimizationFunction func(A [][]float64, y []float64, theta []float64, alpha float64, max_iter int, epsilon float64)  []float64 

// GradientDescent
func GradientDescent(A [][]float64, y []float64, theta []float64, alpha float64, max_iter int, epsilon float64)  []float64 {
	// Initialize theta to 0 vector
	result_theta := theta

	// Get objective function
	J := Objective(A, y, result_theta)

	// while J > epsilon
	iter := 0
	for J > epsilon && iter < max_iter {
		iter += 1
		fmt.Printf("J: %f\n", J)
		// Calculate the gradient of J with respect to theta
		gradient := Gradient(A, y, result_theta)

		// Update theta
		for i := 0; i < len(result_theta); i++ {
			result_theta[i] -= alpha * gradient[i]
		}

		// Update J
		J = Objective(A, y, result_theta)
	}

	return result_theta
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
