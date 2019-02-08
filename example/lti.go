package main

import (
	"fmt"

	"github.com/konimarti/lti"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// define state representation (state-space model)
	system, err := lti.NewSystem(
		mat.NewDense(2, 2, []float64{0, 1, 0, 0}), // A: System matrix
		mat.NewDense(2, 1, []float64{0, 1}),       // B: Control matrix
		mat.NewDense(1, 2, []float64{1, 0}),       // C: Output matrix
		mat.NewDense(1, 1, []float64{0}),          // D: Feedforward matrix
	)
	if err != nil {
		panic(err)
	}

	// check system properties
	fmt.Println("Observable=", system.MustObservable())
	fmt.Println("Controllable=", system.MustControllable())

	// define initial state (x) and control (u) vectors
	x := mat.NewVecDense(2, []float64{0, 1}) // x = (position, velocity)'
	u := mat.NewVecDense(1, []float64{2})    // u = acceleration

	// get derivative vector for new state
	fmt.Println(system.Derivative(x, u))

	// get output vector for new state
	fmt.Println(system.Response(x, u))

	// discretize LTI system and propagate state by time step dt
	dt := 0.1 // dt = time step
	discrete, err := system.Discretize(dt)
	if err != nil {
		panic(err)
	}
	fmt.Println("A_d=", discrete.Ad)
	fmt.Println("B_d=", discrete.Bd)
	fmt.Println("x(k+1)=", discrete.Propagate(x, u))

}
