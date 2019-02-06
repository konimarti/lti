package lti

import "gonum.org/v1/gonum/mat"

type System interface {
	Response(x *mat.VecDense, u *mat.VecDense) *mat.VecDense
}
