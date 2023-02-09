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
	span := 100
	points, values := MakePointsAndValues(size, span)

	// build RBF model
	distance := distances.EuclideanDistance
	kernel := kernels.NewGaussianKernel(distance, map[string]interface{}{"sigma": 0.0005})
	model := rbfinterp.NewRBFInterpolator(kernel, distance)

	// load points
	model.LoadData(points, values)

	// train
	model.Train()

	// create a new point
	x := Noise()*float64(span)
	y := Noise()*float64(span)
	new_point := Point{Dimensionality: 2, Coordinates: []float64{x, y}}
	true_value := HiddenFunction(x, y)

	// predict
	predicted_value, _ := model.Predict(new_point)

	// print results
	fmt.Printf("Point to Query: %v", new_point)
	fmt.Printf("True value: %v\n", true_value)
	fmt.Printf("Predicted value: %v\n", predicted_value)
	fmt.Printf("Error: %v\n", true_value-predicted_value)

	// test exising point from training data
	// (should be exact)
	x = points[0].Coordinates[0]
	y = points[0].Coordinates[1]
	new_point = Point{Dimensionality: 2, Coordinates: []float64{x, y}}
	true_value = values[0]

	// predict
	predicted_value, _ = model.Predict(new_point)

	// print results
	fmt.Printf("True value: %v\n", true_value)
	fmt.Printf("Predicted value: %v\n", predicted_value)
	fmt.Printf("Error: %v\n", true_value-predicted_value)

}

// MakePointsAndValues creates a set of points in 2D space and their corresponding values.
func MakePointsAndValues(span, size int) ([]Point, []float64) {

	points := make([]Point, 0)
	values := make([]float64, 0)
	scale := float64(span)/float64(size)
	for i := 0; i < span; i++ {
		for j := 0; j < span; j++ {
			x := float64(i)*scale - float64(size)/2.0 + Noise()
			y := float64(j)*scale - float64(size)/2.0 + Noise()
			points = append(points, Point{Dimensionality: 2, Coordinates: []float64{x, y}})
			values = append(values, HiddenFunction(x, y)+Noise())
		}
	}

	return points, values
}

func HiddenFunction(x, y float64) float64 {
	return x*x + y*y
}

func Noise() float64 {
	return (2.0*rand.Float64() - 1.0) / 10
}
