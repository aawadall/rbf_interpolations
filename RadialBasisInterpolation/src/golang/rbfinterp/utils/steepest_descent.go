package utils

import "fmt"

// SteepestDescent
func SteepestDescent(A [][]float64, y []float64, theta []float64, alpha float64, max_iter int, epsilon float64) []float64 {
	// Initialize theta to 0 vector
	// theta := make([]float64, len(A[0]))

	fmt.Printf("Steepest Descent Optimization\n")
	// TODO - Implement Steepest Descent Algorithm

	// 1. initialize theta to initial guess
	result_theta := theta

	// 2. calculate gradient using this theta
	gradient := Gradient(A, y, result_theta)
	J := L2Norm(gradient)
	iter := 0
	for J > epsilon && iter < max_iter {
		iter += 1
		// 3. get direction
		direction := SclaerVecMult(-1.0, gradient)

		// 4. get step size
		step_size := LineSearch(A, y, result_theta, direction, alpha)

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

// Min
func Min(a []float64) float64 {
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

// Max
func Max(a []float64) float64 {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
// LineSearch
func LineSearch(A [][]float64, y []float64, theta []float64, direction []float64, alpha float64) float64 {
	// find optimal step size using backtracking line search
	// 1. initialize step size to 1
	step_size := 5.0

	// 2. iterate until stop converging
	for !Converging(A, y, theta, direction, step_size, alpha) {
		// 3. update step size
		step_size *= alpha
	}
	return step_size
}

// Converging
func Converging(A [][]float64, y []float64, theta []float64, direction []float64, step_size float64, alpha float64) bool {
	// 1. calculate theta_new
	theta_new := VecVecAdd(theta, SclaerVecMult(step_size, direction))

	// 2. calculate gradient using theta_new
	gradient_new := Gradient(A, y, theta_new)

	// 3. calculate gradient using theta
	gradient := Gradient(A, y, theta)

	return L2Norm(gradient_new) < L2Norm(gradient)
}
