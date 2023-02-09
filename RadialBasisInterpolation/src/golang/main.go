package main

import (
	"fmt"

	"github.com/aawadall/rbf_interpolations/golang/scratchpad"
)

func main() {
	// Test Gradient Descent
	A := make([][]float64, 3)
	for i := 0; i < 3; i++ {
		A[i] = make([]float64, 3)

		for j := 0; j < 3; j++ {
			A[i][j] = float64(i + j)
		}
	}

	actual_theta := make([]float64, 3)
	actual_theta[0] = 1.0
	actual_theta[1] = 2.0
	actual_theta[2] = 3.0

	actual_y := scratchpad.MatVecDot(A, actual_theta)

	pred_theta := scratchpad.GradientDescent(A, actual_y, 0.0001)

	pred_y := scratchpad.MatVecDot(A, pred_theta)

	for i := 0; i < 3; i++ {
		fmt.Printf("Actual y: %f, Predicted y: %f, err: %f\n", actual_y[i], pred_y[i], actual_y[i]-pred_y[i])
	}
}
