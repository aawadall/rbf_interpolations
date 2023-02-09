package utils

import "math"

// Linear Algebra Utilities
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
	// Gradient g = -2 * A^T * (y - A * theta)
	// Calculate y_hat as A * theta
	y_hat := MatVecDot(A, theta)

	// calculate diff as y - y_hat
	diff := make([]float64, len(y))
	for i := 0; i < len(y); i++ {
		diff[i] = y[i] - y_hat[i]
	}

	// Calculate A transpose
	A_transpose := Transpose(A)

	// Calculate A transpose * diff
	ATdiff := MatVecDot(A_transpose, diff)

	// Calculate -2 * A transpose * diff
	return SclaerVecMult(-2.0, ATdiff)
}

// Transpose
func Transpose(A [][]float64) [][]float64 {
	A_transpose := make([][]float64, len(A[0]))
	for i := 0; i < len(A[0]); i++ {
		A_transpose[i] = make([]float64, len(A))
		for j := 0; j < len(A); j++ {
			A_transpose[i][j] = A[j][i]
		}
	}
	return A_transpose
}

// SclaerVecMult
func SclaerVecMult(scalar float64, vector []float64) []float64 {
	for i := 0; i < len(vector); i++ {
		vector[i] *= scalar
	}
	return vector
}

// DotProduct
func DotProduct(x, y []float64) float64 {
	dot := 0.0
	for i := 0; i < len(x); i++ {
		dot += x[i] * y[i]
	}
	return dot
}

func L2Norm(vector []float64) float64 {
	J := 0.0
	for i := 0; i < len(vector); i++ {
		J += math.Pow(vector[i], 2.0)
	}
	return J
}

// VecVecAdd 
func VecVecAdd(x, y []float64) []float64 {
	z := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		z[i] = x[i] + y[i]
	}
	return z
}