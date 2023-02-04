package distances

// unit tests
import (
	"math"
	"testing"
)

// TestEuclideanDistance
func TestEuclideanDistance(t *testing.T) {
	// happy path
	testCases := []struct {
		x, y Point
		want float64
	}{
		// case 1 (0,0) and (0,0) => 0.0
		{
			x:    Point{Dimensionality: 2, Coordinates: []float64{0.0, 0.0}},
			y:    Point{Dimensionality: 2, Coordinates: []float64{0.0, 0.0}},
			want: 0.0,
		},
		// case 2 (0,0) and (0,1) => 1.0
		{
			x:    Point{Dimensionality: 2, Coordinates: []float64{0.0, 0.0}},
			y:    Point{Dimensionality: 2, Coordinates: []float64{0.0, 1.0}},
			want: 1.0,
		},
		// case 3 (0,0) and (1,0) => 1.0
		{
			x:    Point{Dimensionality: 2, Coordinates: []float64{0.0, 0.0}},
			y:    Point{Dimensionality: 2, Coordinates: []float64{1.0, 0.0}},
			want: 1.0,
		},
		// case 4 (0,0) and (1,1) => sqrt(2)
		{
			x:    Point{Dimensionality: 2, Coordinates: []float64{0.0, 0.0}},
			y:    Point{Dimensionality: 2, Coordinates: []float64{1.0, 1.0}},
			want: math.Sqrt(2.0),
		},
		// case 5 (1,0) and (0,0) => 1.0
		{
			x:    Point{Dimensionality: 2, Coordinates: []float64{1.0, 0.0}},
			y:    Point{Dimensionality: 2, Coordinates: []float64{0.0, 0.0}},
			want: 1.0,
		},
	}

	for _, tc := range testCases {
		got, err := EuclideanDistance(tc.x, tc.y)
		if err != nil {
			t.Errorf("EuclideanDistance(%v, %v) => error: %v", tc.x, tc.y, err)
		}
		if got != tc.want {
			t.Errorf("EuclideanDistance(%v, %v) => %v, want %v", tc.x, tc.y, got, tc.want)
		}
	}
}
