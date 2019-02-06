package main

import (
	"fmt"

	"github.com/konimarti/lti"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// define system type (state-space model)
	system, err := lti.NewDiscrete(
		mat.NewDense(2, 2, []float64{0, 1, 0, 0}), // A: System matrix
		mat.NewDense(2, 1, []float64{0, 1}),       // B: Control matrix
		mat.NewDense(1, 2, []float64{1, 0}),       // C: Output matrix
		mat.NewDense(1, 1, []float64{0}),          // D: Feedforward matrix
	)
	if err != nil {
		panic(err)
	}

	// define initial state (x) and control (u) vectors
	x := mat.NewVecDense(2, []float64{0, 1}) // x = (position, velocity)'
	u := mat.NewVecDense(1, []float64{2})    // u = acceleration
	dt := 0.1                                // dt = time step

	// propagate system by time dt from state x with control u
	nextState, err := system.Propagate(dt, x, u)
	if err != nil {
		panic(err)
	}
	fmt.Println(nextState)

	// get derivative vector for new state
	fmt.Println(system.Derivative(nextState, u))

	// get output vector for new state
	fmt.Println(system.Response(nextState, u))

}
