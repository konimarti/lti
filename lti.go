// Package lti implements a general state-space representation of a linear system.
// A state-space representation can be expressed in a matrix form as
//
// x'(t) = A * x(t) + B * u(t)
// and
// y(t)  = C * x(t) + D * u(t)
//
// where x(t) represents the state and u(t) the control input vectors and the
// matrices are
// A: System matrix,
// B: Control matrix,
// C: Output matrix and
// D: Feedforward matrix
//
package lti

import "gonum.org/v1/gonum/mat"

//LTI represents a general time-continuous state-space LTI system
type LTI interface {
	Observable() (bool, error)
	Controllable() (bool, error)
	Response(x, u *mat.VecDense) *mat.VecDense
}

//DiscreteLTI represents a discretized LTI system
type DiscreteLTI interface {
	LTI
	Propagate(x, u *mat.VecDense) *mat.VecDense
}
