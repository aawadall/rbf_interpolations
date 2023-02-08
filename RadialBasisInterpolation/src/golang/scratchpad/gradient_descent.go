package scratchpad

import "math"

// use this file to experiment with Gradient Descent

// ConjugateGradientDescent
func ConjugateGradientDescent(A [][]float64, y []float64, epsilon float64) []float64 {
	// Given a symmetric positive definite matrix A and a vector y, find a vector x such that Ax = y
	// With tolerance epsilon and initial guess x0 = 0

	// get sizes of A and y to calculate size of x
	m := len(A)    // number of rows in A
	n := len(A[0]) // number of columns in A

	k := len(y) // number of rows in y

	// ensure that A and y are the same size
	if m != k {
		return nil
	}

	// initialize vector x to 0 x n
	x := make([]float64, n)

	// calculate initial gradient g = Ax - y
	g := make([]float64, n)
	g = Add(Dot(A, x), Neg(y))

	// calculate initial residual r = -g
	r := make([]float64, n)
	r = Neg(g)

	// loop until norm(g) < epsilon
	for Norm(g) > epsilon {
		// calclate lambda star as gTr/rTA*r
		lambdaStar := VDot(g, r) / VDot(Dot(A, r), r)
		
		// calculate x = x + lambdaStar * r
		x = Add(x, Scale(lambdaStar, r))

		// calculate g = Ax - y
		g = Add(Dot(A, x), Neg(y))

		// calculate beta = gTAr/rTAr
		beta := VDot(g, Dot(A, r)) / VDot(r, Dot(A, r))

		// calculate r = -g + beta * r
		r = Add(Neg(g), Scale(beta, r))
	}

	return x
}

// Dot
func Dot(A [][]float64, x []float64) []float64 {
	// Given a matrix A and a vector x, return the vector Ax

	// get sizes of A and x to calculate size of Ax
	m := len(A)    // number of rows in A
	n := len(A[0]) // number of columns in A

	k := len(x) // number of rows in x

	// ensure that A and x are the same size
	if n != k {
		return nil
	}

	// initialize vector Ax to 0 x m
	Ax := make([]float64, m)

	// calculate Ax
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			Ax[i] += A[i][j] * x[j]
		}
	}

	return Ax
}

// Add
func Add(x []float64, y []float64) []float64 {
	// Given two vectors x and y, return the vector x + y

	// get sizes of x and y to calculate size of x + y
	m := len(x) // number of rows in x

	k := len(y) // number of rows in y

	// ensure that x and y are the same size
	if m != k {
		return nil
	}

	// initialize vector x + y to 0 x m
	xpy := make([]float64, m)

	// calculate x + y
	for i := 0; i < m; i++ {
		xpy[i] = x[i] + y[i]
	}

	return xpy
}

// Neg
func Neg(x []float64) []float64 {
	// Given a vector x, return the vector -x

	// get size of x to calculate size of -x
	m := len(x) // number of rows in x

	// initialize vector -x to 0 x m
	negx := make([]float64, m)

	// calculate -x
	for i := 0; i < m; i++ {
		negx[i] = -x[i]
	}

	return negx
}

// Norm
func Norm(x []float64) float64 {
	// Given a vector x, return the norm of x

	// get size of x to calculate norm of x
	m := len(x) // number of rows in x

	// initialize norm of x to 0
	normx := 0.0

	// calculate norm of x
	for i := 0; i < m; i++ {
		normx += math.Pow(x[i], 2.0)
	}

	return math.Sqrt(normx)
}


// VDot
func VDot(x []float64, y []float64) float64 {
	// Given two vectors x and y, return the dot product of x and y

	// get sizes of x and y to calculate dot product of x and y
	m := len(x) // number of rows in x

	k := len(y) // number of rows in y

	// ensure that x and y are the same size
	if m != k {
		return 0.0
	}

	// initialize dot product of x and y to 0
	dotxy := 0.0

	// calculate dot product of x and y
	for i := 0; i < m; i++ {
		dotxy += x[i] * y[i]
	}

	return dotxy
}

// Scale
func Scale(lambda float64, x []float64) []float64 {
	// Given a vector x and a scalar lambda, return the vector lambda * x

	// get size of x to calculate size of lambda * x
	m := len(x) // number of rows in x

	// initialize vector lambda * x to 0 x m
	lambdax := make([]float64, m)

	// calculate lambda * x
	for i := 0; i < m; i++ {
		lambdax[i] = lambda * x[i]
	}

	return lambdax
}