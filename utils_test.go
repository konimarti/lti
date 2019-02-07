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
