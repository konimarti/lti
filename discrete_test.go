package lti

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func NewTestDiscrete() (*Discrete, error) {
	sys, _ := NewSystem(
		mat.NewDense(2, 2, []float64{0, 1, 0, 0}), // A
		mat.NewDense(2, 1, []float64{0, 1}),       // B
		mat.NewDense(1, 2, []float64{1, 0}),       // C
		mat.NewDense(1, 1, []float64{0}),          // D
	)
	dt := 0.1
	//return NewDiscrete(sys.A, sys.B, nil, nil, dt)
	return sys.Discretize(dt)
}

func TestPropagate(t *testing.T) {
	sys, err := NewTestDiscrete()
	if err != nil {
		t.Error("Internal error in creating test system")
	}

	state := mat.NewVecDense(2, []float64{0, 1}) // x = position, velocity
	input := mat.NewVecDense(1, []float64{2})    // u = accelartion
	newState := sys.Propagate(state, input)
	if err != nil {
		fmt.Println(err)
		t.Error("Propagate returned error")
	}

	//
	expectedState := mat.NewVecDense(2, []float64{0.11, 1.2})
	if !mat.EqualApprox(newState, expectedState, 1e-4) {
		fmt.Println("Returned:", newState)
		fmt.Println("Expected:", expectedState)
		t.Error("Propagate returned wrong state")
	}
}
