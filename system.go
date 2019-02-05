package lti

import "gonum.org/v1/gonum/mat"

type System interface {
	Propagate(dt float64, x *mat.VecDense, u *mat.VecDense) *mat.VecDense
	Response(x *mat.VecDense, u *mat.VecDense) *mat.VecDense
}
