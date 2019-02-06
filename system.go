package lti

import "gonum.org/v1/gonum/mat"

//System represents the state-space representation of LTI system
type System interface {
	Response(x *mat.VecDense, u *mat.VecDense) *mat.VecDense
}
