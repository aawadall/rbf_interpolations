package main

import (
	"fmt"
	"math"

	"github.com/aawadall/rbf_interpolations/golang/scratchpad"
)

func main() {
	// HACK - test scratchpad.ConjugateGradientDescent on A and y and test
	A := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
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
	error := 0.0
	for i := 0; i < 3; i++ {
		error += math.Pow(y_hat[i]-y[i], 2.0)
	}

	fmt.Printf("error: %f\n", error)

	// TODO - implement
}
