package main

import (
	"fmt"
	"math/rand"

	"github.com/aawadall/rbf_interpolations/golang/rbfinterp"
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/distances"
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/kernels"
	"github.com/aawadall/rbf_interpolations/golang/rbfinterp/types"
)

// define types
type Point = types.Point

func main() {
	// let us build a scenario where we have a set of points in 2D space and their corresponding values
	size := 10
	points, values := MakePointsAndValues(size)

	// build RBF model
	distance := distances.EuclideanDistance
	kernel := kernels.NewGaussianKernel(distance, map[string]interface{}{"sigma": 0.1})
	model := rbfinterp.NewRBFInterpolator(kernel, distance)

	// load points
	model.LoadData(points, values)

	// train
	model.Train()

	// create a new point
	x := 10.0
	y := 11.0
	new_point := Point{Dimensionality: 2, Coordinates: []float64{x, y}}
	true_value := HiddenFunction(x, y)

	// predict
	predicted_value, _ := model.Predict(new_point)

	// print results
	fmt.Printf("True value: %v\n", true_value)
	fmt.Printf("Predicted value: %v\n", predicted_value)
	fmt.Printf("Error: %v\n", true_value-predicted_value)
}

// MakePointsAndValues creates a set of points in 2D space and their corresponding values.
func MakePointsAndValues(size int) ([]Point, []float64) {

	points := make([]Point, 0)
	values := make([]float64, 0)

	for i := 0; i < size; i++ {
		x := (rand.Float64() * 100.0) - 50.0
		y := (rand.Float64() * 100.0) - 50.0
		points = append(points, Point{Dimensionality: 2, Coordinates: []float64{x, y}})
		values = append(values, HiddenFunction(x, y)+Noise())
	}

	return points, values
}

func HiddenFunction(x, y float64) float64 {
	return x*x + y*y
}

func Noise() float64 {
	return (2.0*rand.Float64() - 1.0) / 10
}
