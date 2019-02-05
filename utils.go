package lti

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

//Discretize
func Discretize(m *mat.Dense, dt float64) (*mat.Dense, error) {
	// m_d = exp(m * dt)

	// check if matrix m is square
	if r, c := m.Dims(); r != c {
		return m, errors.New("Discretize: matrix is not square")
	}

	// tmp = m * dt
	var tmp mat.Dense
	tmp.Scale(dt, m)

	// exp( tmp )
	var md mat.Dense
	md.Exp(&tmp)

	return &md, nil
}

// Integrate
// Source: https://math.stackexchange.com/questions/658276/integral-of-matrix-exponential
// Int_0^T exp(A * t) * B dt = T [ exp(AT) - AT ] * B
func Integrate(a *mat.Dense, b *mat.Dense, dt float64) (*mat.Dense, error) {
	// exp(A t)
	ad, err := Discretize(a, dt)
	if err != nil {
		return nil, errors.New("discretization failed")
	}

	// (At)
	var at mat.Dense
	at.Scale(dt, a)

	// exp(A t) - At
	var diff mat.Dense
	diff.Sub(ad, &at)

	// diff * B
	var bd mat.Dense
	bd.Mul(&diff, b)

	bd.Scale(dt, &bd)

	return &bd, nil
}
