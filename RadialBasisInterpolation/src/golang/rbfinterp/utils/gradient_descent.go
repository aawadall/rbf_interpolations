package utils

import (
	"fmt"
)

type OptimizationFunction func(A [][]float64, y []float64, theta []float64, configuration map[string]interface{})  []float64 

// GradientDescent
func GradientDescent(A [][]float64, y []float64, theta []float64, configuration map[string]interface{})  []float64 {
	fmt.Printf("Gradient Descent Optimization\n")
	// Lookup Configuration
	
	// Alpha 
	alpha := 0.1 
	if alphaParam, ok := configuration["alpha"]; ok {
		alpha = alphaParam.(float64)
	}
	
	// Max Iteration
	maxIter := 1000
	if maxIterParam, ok := configuration["maxIter"]; ok {
		maxIter = maxIterParam.(int)
	}

	// Epsilon 
	epsilon := 0.0001
	if epsilonParam, ok := configuration["epsilon"]; ok {
		epsilon = epsilonParam.(float64)
	}


	// Print Configuration
	fmt.Printf("Configuration: alpha: %f, maxIter: %d, epsilon: %f\n", alpha, maxIter, epsilon)
	
	// Initialize theta to 0 vector
	resultTheta := theta

	// Get objective function
	J := Objective(A, y, resultTheta)

	// while J > epsilon
	iter := 0
	for J > epsilon && iter < maxIter {
		iter += 1
		// Calculate the gradient of J with respect to theta
		gradient := Gradient(A, y, resultTheta)
		
		// Update theta
		for i := 0; i < len(resultTheta); i++ {
			resultTheta[i] -= alpha * gradient[i]
		}
		
		// Update J
		old_j := J
		J = Objective(A, y, resultTheta)
		fmt.Printf("[%d] \tJ: %f -> improved %0.2f%%\n", iter, J, (old_j-J)/old_j*100.0)
	}

	return resultTheta
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
