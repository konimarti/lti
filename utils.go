package lti

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

//Discretize
func Discretize(m *mat.Dense, t float64) (*mat.Dense, error) {
	// m_d = exp(m * t)

	// check if matrix m is square
	if r, c := m.Dims(); r != c {
		return m, errors.New("Discretize: matrix is not square")
	}

	// tmp = m * t
	var tmp mat.Dense
	tmp.Scale(t, m)

	// exp( tmp )
	var md mat.Dense
	md.Exp(&tmp)

	return &md, nil
}

// Integrate
// Source: https://math.stackexchange.com/questions/658276/integral-of-matrix-exponential
// Int_0^T exp(A t) B dt = T [ exp(AT) - AT ] * B
func Integrate(ad *mat.Dense, a *mat.Dense, b *mat.Dense, t float64) (*mat.Dense, error) {

	// (At)
	var at mat.Dense
	at.Scale(t, a)

	// exp(A t) - At
	var diff mat.Dense
	diff.Sub(ad, &at)

	// diff * B
	var bd mat.Dense
	bd.Mul(&diff, b)

	bd.Scale(t, &bd)

	return &bd, nil
}
