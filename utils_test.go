package lti

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestDiscretize(t *testing.T) {
	dt := 0.1

	// check for non square matrices
	notSquare := mat.NewDense(4, 5, nil)
	_, err := discretize(notSquare, dt)
	if err == nil {
		t.Error("Should have returned an error")
	}

	// check for correct calculation
	m := mat.NewDense(2, 2, []float64{1, 0, 0, 1})
	md, err := discretize(m, dt)
	correct := mat.NewDense(2, 2, []float64{1.1052, 0.0, 0.0, 1.1052})
	if !mat.EqualApprox(md, correct, 1e-4) {
		t.Error("Discretize does not return correct matrix")
	}

}
