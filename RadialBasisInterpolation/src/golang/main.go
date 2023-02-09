package main

import (
	"fmt"
	"math"

	"github.com/aawadall/rbf_interpolations/golang/scratchpad"
)

func main() {
	// HACK - test scratchpad.ConjugateGradientDescent on A and y and test
	A := [][]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	y := []float64{1, 2, 3}

	x := scratchpad.ConjugateGradientDescent(A, y, 0.001)

	// test error
	y_hat := make([]float64, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			y_hat[i] += A[i][j] * x[j]
		}
	}

	// test error
	
	for i := 0; i < 3; i++ {
		fmt.Printf("y_hat[%d] = %f\n", i, y_hat[i])
		fmt.Printf("y[%d] = %f\n", i, y[i])
		fmt.Printf("error = %f\n", math.Abs(y_hat[i] - y[i]))		
	}


	// TODO - implement
}
