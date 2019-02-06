package lti

import "gonum.org/v1/gonum/mat"

// The System interface represents the state-space representation of LTI systems.
type System interface {
	Response(x *mat.VecDense, u *mat.VecDense) *mat.VecDense
}
