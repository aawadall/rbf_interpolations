package utils

import "math"

// Linear Algebra Utilities
func MatVecDot(A [][]float64, theta []float64) []float64 {
	// y_hat := make([]float64, len(A))
	// for i := 0; i < len(A); i++ {
	// 	for j := 0; j < len(A[0]); j++ {
	// 		y_hat[i] += A[i][j] * theta[j]
	// 	}
	// }
	// return y_hat
	return MatVecDotGoRoutine(A, theta)
}

// MatVecDotGoRoutine
func MatVecDotGoRoutine(A [][]float64, theta []float64) []float64 {
	y_hat := make([]float64, len(A))
	for i := 0; i < len(A); i++ {
		go func(i int) {
			for j := 0; j < len(A[0]); j++ {
				y_hat[i] += A[i][j] * theta[j]
			}
		}(i)
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
	diff := VecVecAdd(y, SclaerVecMult(-1.0, y_hat))
	
	// Calculate A transpose
	A_transpose := Transpose(A)

	// Calculate A transpose * diff
	ATdiff := MatVecDot(A_transpose, diff)

	// Calculate -2 * A transpose * diff
	return SclaerVecMult(-2.0, ATdiff)
}

// Transpose
func Transpose(A [][]float64) [][]float64 {
	// A_transpose := make([][]float64, len(A[0]))
	// for i := 0; i < len(A[0]); i++ {
	// 	A_transpose[i] = make([]float64, len(A))
	// 	for j := 0; j < len(A); j++ {
	// 		A_transpose[i][j] = A[j][i]
	// 	}
	// }
	// return A_transpose
	return TransposeGoRoutine(A)
}

// TransposeGoRoutine 
func TransposeGoRoutine(A [][]float64) [][]float64 {
	A_transpose := make([][]float64, len(A[0]))
	for i := 0; i < len(A[0]); i++ {
		A_transpose[i] = make([]float64, len(A))
	}

	for i := 0; i < len(A[0]); i++ {
		go func(i int) {
			for j := 0; j < len(A); j++ {
				A_transpose[i][j] = A[j][i]
			}
		}(i)
	}
	return A_transpose
}
// SclaerVecMult
func SclaerVecMult(scalar float64, vector []float64) []float64 {
	// for i := 0; i < len(vector); i++ {
	// 	vector[i] *= scalar
	// }
	// return vector
	return ScalarVecMultGoRoutine(scalar, vector)
}

// ScalarVecMultGoRoutine
func ScalarVecMultGoRoutine(scalar float64, vector []float64) []float64 {
	for i := 0; i < len(vector); i++ {
		go func(i int) {
			vector[i] *= scalar
		}(i)
	}
	return vector
}
// DotProduct
func DotProduct(x, y []float64) float64 {
	// dot := 0.0
	// for i := 0; i < len(x); i++ {
	// 	dot += x[i] * y[i]
	// }
	// return dot
	return DotProductGoRoutine(x, y)
}

// DotProductGoRoutine 
func DotProductGoRoutine(x, y []float64) float64 {
	dot := 0.0
	for i := 0; i < len(x); i++ {
		go func(i int) {
			dot += x[i] * y[i]
		}(i)
	}
	return dot
}

func L2Norm(vector []float64) float64 {
	// J := 0.0
	// for i := 0; i < len(vector); i++ {
	// 	J += math.Pow(vector[i], 2.0)
	// }
	// return J
	return L2NormGoRoutine(vector)
}

// L2NormGoRoutine
func L2NormGoRoutine(vector []float64) float64 {
	J := 0.0
	for i := 0; i < len(vector); i++ {
		go func(i int) {
			J += math.Pow(vector[i], 2.0)
		}(i)
	}
	return math.Pow(J, 0.5)
}

// VecVecAdd 
func VecVecAdd(x, y []float64) []float64 {
	// z := make([]float64, len(x))
	// for i := 0; i < len(x); i++ {
	// 	z[i] = x[i] + y[i]
	// }
	// return z
	return VecVecAddGoRoutine(x, y)
}

// VecVecAddGoRoutine
func VecVecAddGoRoutine(x, y []float64) []float64 {
	z := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		go func(i int) {
			z[i] = x[i] + y[i]
		}(i)
	}
	return z
}

// MatrixNorm
func MatrixNorm(A [][]float64) float64 {
	// J := 0.0
	// for i := 0; i < len(A); i++ {
	// 	for j := 0; j < len(A[0]); j++ {
	// 		J += math.Pow(A[i][j], 2.0)
	// 	}
	// }
	// return J
	return MatrixNormGoRoutine(A)
}

// MatrixNormGoRoutine
func MatrixNormGoRoutine(A [][]float64) float64 {
	J := 0.0
	for i := 0; i < len(A); i++ {
		go func(i int) {
			for j := 0; j < len(A[0]); j++ {
				J += math.Pow(A[i][j], 2.0)
			}
		}(i)
	}
	return math.Pow(J, 0.5)
}

// MatrixScale 
func MatrixScale(A [][]float64, scalar float64) [][]float64 {
	// for i := 0; i < len(A); i++ {
	// 	for j := 0; j < len(A[0]); j++ {
	// 		A[i][j] *= scalar
	// 	}
	// }
	// return A
	return MatrixScaleGoRoutine(A, scalar)
}

// MatrixScaleGoRoutine
func MatrixScaleGoRoutine(A [][]float64, scalar float64) [][]float64 {
	for i := 0; i < len(A); i++ {
		go func(i int) {
			for j := 0; j < len(A[0]); j++ {
				A[i][j] *= scalar
			}
		}(i)
	}
	return A
}

// VectorScale
func VectorScale(x []float64, scalar float64) []float64 {
	// for i := 0; i < len(x); i++ {
	// 	x[i] *= scalar
	// }
	// return x
	return VectorScaleGoRoutine(x, scalar)
}

// VectorScaleGoRoutine
func VectorScaleGoRoutine(x []float64, scalar float64) []float64 {
	for i := 0; i < len(x); i++ {
		go func(i int) {
			x[i] *= scalar
		}(i)
	}
	return x
}