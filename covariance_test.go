package lti

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestCovariancePredict(t *testing.T) {

	md := mat.NewDense(3, 3, []float64{
		0, 1, 0,
		0, 0, 1,
		1, 0, 0,
	})
	systemNoise := NewCovariance(md)

	p := mat.NewDense(3, 3, []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	})
	pNext := systemNoise.Predict(p, nil)

	expected1 := mat.NewDense(3, 3, []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	})
	if !mat.EqualApprox(pNext, expected1, 1e-8) {
		fmt.Println("received:", pNext)
		fmt.Println("expected:", expected1)
		t.Error("predict failed without noise")
	}

	// test with extra noise
	noise := mat.NewDense(3, 3, []float64{
		0.1, 0.1, 0.3,
		0.1, 0.2, 0.1,
		0.3, 0.1, 0.1,
	})
	pNext = systemNoise.Predict(p, noise)

	expected2 := mat.NewDense(3, 3, []float64{
		1.1, 0.1, 0.3,
		0.1, 1.2, 0.1,
		0.3, 0.1, 1.1,
	})
	if !mat.EqualApprox(pNext, expected2, 1e-8) {
		fmt.Println("received:", pNext)
		fmt.Println("expected:", expected2)
		t.Error("predict failed with noise")
	}

}
