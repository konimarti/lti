package lti

import (
	"fmt"
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
	md, _ := discretize(m, dt)
	correct := mat.NewDense(2, 2, []float64{1.1052, 0.0, 0.0, 1.1052})
	if !mat.EqualApprox(md, correct, 1e-4) {
		t.Error("Discretize does not return correct matrix")
	}
}

func TestIntegrate(t *testing.T) {
	dt := 0.1

	a := mat.NewDense(3, 3, []float64{
		0, 1, 0,
		0, 0, 1,
		0, 0, 0,
	})
	b := mat.NewDense(3, 1, []float64{
		0,
		0,
		1,
	})

	bd, err := integrate(a, b, dt)
	if err != nil {
		t.Error("failed to integrate")
	}

	correct := mat.NewDense(3, 1, []float64{
		1.0 / 6.0 * dt * dt * dt,
		0.5 * dt * dt,
		dt,
	})
	if !mat.EqualApprox(bd, correct, 1e-4) {
		fmt.Println("received:", bd)
		fmt.Println("expected:", correct)
		t.Error("Integrate does not return correct matrix")
	}
}

func TestRank(t *testing.T) {
	var config = []struct {
		M    *mat.Dense
		Rank int
	}{
		{
			M:    mat.NewDense(2, 2, []float64{0, 1, 0, 0}),
			Rank: 1,
		},
		{
			M:    mat.NewDense(2, 2, []float64{1, 0, 0, 1}),
			Rank: 2,
		},
		{
			M: mat.NewDense(3, 3, []float64{
				1, 2, 3,
				0, 4, 6,
				0, 0, 9,
			}),
			Rank: 3,
		},
		{
			M: mat.NewDense(3, 3, []float64{
				1, 2, 3,
				0, 4, 6,
				0, 6, 9,
			}),
			Rank: 2,
		},
		{
			M: mat.NewDense(3, 3, []float64{
				1, 2, 3,
				2, 4, 6,
				3, 6, 9,
			}),
			Rank: 1,
		},
		{
			M: mat.NewDense(4, 2, []float64{
				1, 2,
				0, 4,
				3, 0,
				4, 0,
			}),
			Rank: 2,
		},
		{
			M: mat.NewDense(4, 2, []float64{
				0, 2,
				0, 4,
				3, 0,
				4, 0,
			}),
			Rank: 2,
		},
		{
			M: mat.NewDense(4, 2, []float64{
				1, 2,
				2, 4,
				3, 6,
				4, 8,
			}),
			Rank: 1,
		},
	}

	for _, cfg := range config {
		if r, _ := rank(cfg.M); r != cfg.Rank {
			fmt.Println("m=", cfg.M)
			fmt.Println("expected:", cfg.Rank)
			fmt.Println("received:", r)
			t.Error("rank failed")
		}
	}
}

func TestCheckControllability(t *testing.T) {

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

	for _, cfg := range config {
		if ok, _ := checkControllability(cfg.A, cfg.B); ok != cfg.Want {
			fmt.Println("A=", cfg.A)
			fmt.Println("B=", cfg.B)
			fmt.Println("received:", ok)
			fmt.Println("expected:", cfg.Want)
			t.Error("check controllability failed")
		}
	}

}

func TestCheckObservability(t *testing.T) {

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

	for _, cfg := range config {
		if ok, _ := checkObservability(cfg.A, cfg.C); ok != cfg.Want {
			fmt.Println("A=", cfg.A)
			fmt.Println("C=", cfg.C)
			fmt.Println("received:", ok)
			fmt.Println("expected:", cfg.Want)
			t.Error("observable failed")
		}
	}

}

func TestMultAndSumOp(t *testing.T) {

	var config = []struct {
		A    *mat.Dense
		X    *mat.VecDense
		B    *mat.Dense
		U    *mat.VecDense
		Want *mat.VecDense
	}{
		{
			A: mat.NewDense(2, 2, []float64{
				0, 1,
				0, 0,
			}),
			X: mat.NewVecDense(2, []float64{
				0, 1,
			}),
			B: mat.NewDense(2, 2, []float64{
				1, 0,
				0, 1,
			}),
			U: mat.NewVecDense(2, []float64{
				1, 0,
			}),
			Want: mat.NewVecDense(2, []float64{
				2, 0,
			}),
		},
		{
			A: mat.NewDense(3, 3, []float64{
				0, 1, 0,
				0, 0, 1,
				1, 0, 0,
			}),
			X: mat.NewVecDense(3, []float64{
				1, 2, 3,
			}),
			B: mat.NewDense(3, 2, []float64{
				0, 1,
				1, 0,
				0, 0,
			}),
			U: mat.NewVecDense(2, []float64{
				3, 2,
			}),
			Want: mat.NewVecDense(3, []float64{
				4, 6, 1,
			}),
		},
	}

	for _, cfg := range config {
		if sum := multAndSumOp(cfg.A, cfg.X, cfg.B, cfg.U); !mat.EqualApprox(sum, cfg.Want, 1e-6) {
			fmt.Println("received:", sum)
			fmt.Println("expected:", cfg.Want)
			t.Error("Multiplication and Summation operation failed")
		}
	}

}
