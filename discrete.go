package lti

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

// Discrete represents a discrete LTI system.
//
// The parameters are:
// 	A_d: Discretized Ssystem matrix
// 	B_d: Discretized Control matrix
//
//
type Discrete struct {
	Ad *mat.Dense
	Bd *mat.Dense
}

//NewDiscrete returns a Discrete struct
func NewDiscrete(A, B *mat.Dense, dt float64) (*Discrete, error) {

	// A_d = exp(A*dt)
	ad, err := discretize(A, dt)
	if err != nil {
		return nil, errors.New("discretization of A failed")
	}

	// B_d = Int_0^T exp(A*dt) * B dt
	bd, err := integrate(A, B, dt)
	if err != nil {
		return nil, errors.New("discretization of B failed")
	}

	return &Discrete{
		Ad: ad,
		Bd: bd,
	}, nil
}

// Propagate state x(k) by a time step dt to x(k+1) = A_discretized * x + B_discretized * u
func (d *Discrete) Propagate(x *mat.VecDense, u *mat.VecDense) *mat.VecDense {

	// x(k+1) = A_d * x + B_d * u
	var xk1, adx, bdu mat.VecDense
	adx.MulVec(d.Ad, x)
	bdu.MulVec(d.Bd, u)
	xk1.AddVec(&adx, &bdu)

	return &xk1
}
