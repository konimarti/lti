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
	Ad          *mat.Dense
	Bd          *mat.Dense
	C           *mat.Dense
	D           *mat.Dense
	ax, bu, sum mat.VecDense // Workspace for multAndSumOp
}

//NewDiscrete returns a Discrete struct
func NewDiscrete(A, B, C, D *mat.Dense, dt float64) (*Discrete, error) {

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
		C:  C,
		D:  D,
	}, nil
}

// Predict predicts  x(k+1) = A_discretized * x(k) + B_discretized * u(k)
func (d *Discrete) Predict(x *mat.VecDense, u *mat.VecDense) *mat.VecDense {
	// x(k+1) = A_d * x + B_d * u

	return multAndSumOp(d.Ad, x, d.Bd, u, d.ax, d.bu, d.sum)
}

//Response returns the output vector y(t) = C * x(t) + D * u(t)
func (d *Discrete) Response(x *mat.VecDense, u *mat.VecDense) *mat.VecDense {
	// y(t) = C * x(t) + D * u(t)
	return multAndSumOp(d.C, x, d.D, u, d.ax, d.bu, d.sum)
}

// Controllable checks the controllability of the LTI system.
func (d *Discrete) Controllable() (bool, error) {
	// system is controllable if
	// rank( [B, A B, A^2 B, A^n-1 B] ) = n
	return checkControllability(d.Ad, d.Bd)
}

// Observable checks the observability of the LTI system.
func (d *Discrete) Observable() (bool, error) {
	// system is observable if
	// rank( S=[C, C A, C A^2, ..., C A^n-1]' ) = n
	return checkObservability(d.Ad, d.C)
}
