package lti

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestIdealDiscretization(t *testing.T) {
	dt := 0.1
	var config = []struct {
		A    *mat.Dense
		T    float64
		M    *mat.Dense
		Want *mat.Dense
	}{
		{
			A: mat.NewDense(2, 2, []float64{
				0, 1,
				0, 0,
			}),
			T: dt,
			M: nil,
			Want: mat.NewDense(2, 2, []float64{
				1, dt,
				0, 1,
			}),
		},
		{
			A: mat.NewDense(2, 2, []float64{
				0, 1,
				0, 0,
			}),
			T: dt,
			M: mat.NewDense(2, 2, []float64{
				0, 1,
				1, 0,
			}),
			Want: mat.NewDense(2, 2, []float64{
				dt, 1,
				1, 0,
			}),
		},
	}

	for _, cfg := range config {
		got, err := IdealDiscretization(cfg.A, cfg.T, cfg.M)
		if err == nil {
			if !mat.EqualApprox(got, cfg.Want, 1e-8) {
				fmt.Println("received=", got)
				fmt.Println("expected=", cfg.Want)
				t.Error("ideal discretization returned wrong result")
			}
		} else {
			fmt.Println(err)
			t.Error("error received in ideal discretization")
		}

	}

}

func TestRealDiscretization(t *testing.T) {
	dt := 0.1
	var config = []struct {
		A    *mat.Dense
		T    float64
		M    *mat.Dense
		Want *mat.Dense
	}{
		{
			A: mat.NewDense(2, 2, []float64{
				0, 1,
				0, 0,
			}),
			T: dt,
			M: mat.NewDense(2, 1, []float64{
				0,
				1,
			}),
			Want: mat.NewDense(2, 1, []float64{
				0.5 * dt * dt,
				dt,
			}),
		},
	}

	for _, cfg := range config {
		got, err := RealDiscretization(cfg.A, cfg.T, cfg.M)
		if err == nil {
			if !mat.EqualApprox(got, cfg.Want, 1e-8) {
				fmt.Println("received=", got)
				fmt.Println("expected=", cfg.Want)
				t.Error("ideal discretization returned wrong result")
			}
		} else {
			fmt.Println(err)
			t.Error("error received in ideal discretization")
		}

	}

}
