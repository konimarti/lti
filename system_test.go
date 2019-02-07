package lti

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func NewTestSystem() (*System, error) {
	return NewSystem(
		mat.NewDense(2, 2, []float64{0, 1, 0, 0}), // A
		mat.NewDense(2, 1, []float64{0, 1}),       // B
		mat.NewDense(1, 2, []float64{1, 0}),       // C
		mat.NewDense(1, 1, []float64{0}),          // D
	)
}

func TestResponse(t *testing.T) {
	sys, err := NewTestSystem()
	if err != nil {
		t.Error("Internal error in creating test system")
	}

	state := mat.NewVecDense(2, []float64{1.1, 1}) // x = position, velocity
	input := mat.NewVecDense(1, []float64{2})      // u = accelartion
	response := sys.Response(state, input)

	//
	expected := mat.NewVecDense(1, []float64{1.1})
	if !mat.EqualApprox(response, expected, 1e-4) {
		fmt.Println("Returned:", response)
		fmt.Println("Expected:", expected)
		t.Error("Response returned wrong state")
	}
}

func TestDerivative(t *testing.T) {
	sys, err := NewTestSystem()
	if err != nil {
		t.Error("Internal error in creating test system")
	}

	state := mat.NewVecDense(2, []float64{1.1, 1}) // x = position, velocity
	input := mat.NewVecDense(1, []float64{2})      // u = accelartion
	deriv := sys.Derivative(state, input)

	//
	expected := mat.NewVecDense(2, []float64{1, 2})
	if !mat.EqualApprox(deriv, expected, 1e-4) {
		fmt.Println("Returned:", deriv)
		fmt.Println("Expected:", expected)
		t.Error("Response returned wrong state")
	}

}

func TestControllable(t *testing.T) {

	var config = []struct {
		A    *mat.Dense
		B    *mat.Dense
		Want bool
	}{
		{
			A: mat.NewDense(3, 3, []float64{
				0, 1, 0,
				0, 0, 1,
				0, 0, 0}),
			B: mat.NewDense(3, 1, []float64{
				0, 0, 0}),
			Want: false,
		},
		{
			A: mat.NewDense(2, 2, []float64{
				0, 1,
				0, 0}),
			B: mat.NewDense(2, 1, []float64{
				0, 1}),
			Want: true,
		},
	}

	sys := System{}

	for _, cfg := range config {
		sys.A = cfg.A
		sys.B = cfg.B
		if ok, _ := sys.Controllable(); ok != cfg.Want || ok != sys.MustControllable() {
			fmt.Println("A=", cfg.A)
			fmt.Println("B=", cfg.B)
			fmt.Println("received:", ok)
			fmt.Println("expected:", cfg.Want)
			t.Error("controllable failed")
		}
	}

}

func TestObservable(t *testing.T) {

	var config = []struct {
		A    *mat.Dense
		C    *mat.Dense
		Want bool
	}{
		{
			A: mat.NewDense(2, 2, []float64{
				0, 1,
				0, 0}),
			C: mat.NewDense(1, 2, []float64{
				1, 0}),
			Want: true,
		},
		{
			A: mat.NewDense(3, 3, []float64{
				0, 1, 0,
				0, 0, 1,
				0, 0, 0}),
			C: mat.NewDense(1, 3, []float64{
				0, 1, 0}),
			Want: false,
		},
	}

	sys := System{}

	for _, cfg := range config {
		sys.A = cfg.A
		sys.C = cfg.C
		if ok, _ := sys.Observable(); ok != cfg.Want || ok != sys.MustObservable() {
			fmt.Println("A=", cfg.A)
			fmt.Println("C=", cfg.C)
			fmt.Println("received:", ok)
			fmt.Println("expected:", cfg.Want)
			t.Error("observable failed")
		}
	}

}
