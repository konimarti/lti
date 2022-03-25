package lti

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

// System represents the state equations of time-continuous, linear systems
//
// The parameters are:
// 	A: System matrix
// 	B: Control matrix
// 	C: Output matrix
// 	D: Feedforward matrix
//
//
type System struct {
	A           *mat.Dense
	B           *mat.Dense
	C           *mat.Dense
	D           *mat.Dense
	ax, bu, sum mat.VecDense // Workspace for multAndSumOp
}

//NewSystem returns a System struct and checks the matrix dimensions
func NewSystem(A, B, C, D *mat.Dense) (*System, error) {

	// A (n x n)
	ar, ac := A.Dims()
	if ar != ac {
		return nil, errors.New("A should be squared")
	}
	// B (n x k)
	br, bc := B.Dims()
	if br != ar {
		return nil, errors.New("B row should be equal to A row dim")
	}

	// C (l x n)
	cr, cc := C.Dims()
	if cc != ar {
		return nil, errors.New("C col should be equal to A row dim")
	}

	// D (l x k)
	dr, dc := D.Dims()
	if dr != cr {
		return nil, errors.New("D row should be equal to C row dim")
	}
	if dc != bc {
		return nil, errors.New("D col should be equal to B col dim")
	}

	return &System{
		A: A,
		B: B,
		C: C,
		D: D,
	}, nil
}

//Derivative returns the derivative vetor x'(t) = A * x(t) + B * u(t)
func (s *System) Derivative(x, u *mat.VecDense) *mat.VecDense {
	// x'(t) = A * x(t) + B * u(t)
	return multAndSumOp(s.A, x, s.B, u, s.ax, s.bu, s.sum)
}

//Response returns the output vector y(t) = C * x(t) + D * u(t)
func (s *System) Response(x *mat.VecDense, u *mat.VecDense) *mat.VecDense {
	// y(t) = C * x(t) + D * u(t)
	return multAndSumOp(s.C, x, s.D, u, s.ax, s.bu, s.sum)
}

// MustControllable checks the controllability
// and panics when error occurs
func (s *System) MustControllable() bool {
	ok, err := s.Controllable()
	if err != nil {
		panic(err)
	}
	return ok
}

// Controllable checks the controllability of the LTI system.
func (s *System) Controllable() (bool, error) {
	// system is controllable if
	// rank( [B, A B, A^2 B, A^n-1 B] ) = n
	return checkControllability(s.A, s.B)
}

// MustObservable checks the observability of the LTI system.
// and panics when error occurs
func (s *System) MustObservable() bool {
	ok, err := s.Observable()
	if err != nil {
		panic(err)
	}
	return ok
}

// Observable checks the observability of the LTI system.
func (s *System) Observable() (bool, error) {
	// system is observable if
	// rank( S=[C, C A, C A^2, ..., C A^n-1]' ) = n
	return checkObservability(s.A, s.C)
}

// Discretize discretizes the time-continuous LTI into an explicit time-discrete LTI system
func (s *System) Discretize(dt float64) (*Discrete, error) {
	return NewDiscrete(s.A, s.B, s.C, s.D, dt)
}
