package utils

import "fmt"

// SteepestDescent
func RegularizedSteepestDescent(A [][]float64, y []float64, theta []float64, configuration map[string]interface{}) []float64 {
	// Initialize theta to 0 vector
	// theta := make([]float64, len(A[0]))

	fmt.Printf("Steepest Descent Optimization\n")
	// TODO - Implement Steepest Descent Algorithm

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

	// Lambda 2
	lambda2 := 0.1
	if lambda2PAram, ok := configuration["lambda2"]; ok {
		
		lambda2 = lambda2PAram.(float64)
		
	}

	// Print Configuration
	fmt.Printf("Configuration: alpha: %f, maxIter: %d, epsilon: %f, lambda2: %f\n", alpha, maxIter, epsilon, lambda2)
	// 1. initialize theta to initial guess
	result_theta := theta

	// 2. calculate gradient using this theta
	gradient := Gradient(A, y, result_theta)
	J := L2Norm(gradient) 
	iter := 0
	for J > epsilon  && iter < maxIter {
		iter += 1
		// 3. get direction
		direction := SclaerVecMult(-1.0, gradient)

		// 4. get step size
		step_size := LineSearch(A, y, result_theta, direction, alpha, lambda2)

		// 5. update theta
		result_theta = VecVecAdd(result_theta, SclaerVecMult(step_size, direction))

		// 6. calculate gradient using this theta
		gradient = Gradient(A, y, result_theta)
		old_j := J
		J = L2Norm(gradient) 
		minTheta := Min(result_theta)
		maxTheta := Max(result_theta)
		fmt.Printf("[%d] \tJ: %f -> step size %f, improved %0.2f%% theta: max[%f], min[%f]\n", iter, J, step_size, (old_j-J)/old_j*100.0, maxTheta, minTheta)
	}
	return result_theta
}
