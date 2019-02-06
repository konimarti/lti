package lti

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

//Discrete represents a discrete LTI system
type Discrete struct {
	A *mat.Dense
	B *mat.Dense
	C *mat.Dense
	D *mat.Dense
}

//NewDiscrete returns a Discrete struct and checks the matrix dimensions
func NewDiscrete(A, B, C, D *mat.Dense) (*Discrete, error) {

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

	return &Discrete{
		A: A,
		B: B,
		C: C,
		D: D,
	}, nil
}

//Derivative returns the derivative vetor x'(t) = A * x(t) + B * u(t)
func (d *Discrete) Derivative(x, u *mat.VecDense) *mat.VecDense {
	// x'(t) = A * x(t) + B * u(t)

	// A * x(t)
	var ax mat.VecDense
	ax.MulVec(d.A, x)

	// B * u(t)
	var bx mat.VecDense
	bx.MulVec(d.B, u)

	// xderiv = A x(t) + B u(t)
	var xderiv mat.VecDense
	xderiv.AddVec(&ax, &bx)

	return &xderiv
}

// Propagate state x(k) by a time step dt to x(k+1) = A_discretized * x + B_discretized * u
func (d *Discrete) Propagate(dt float64, x *mat.VecDense, u *mat.VecDense) (*mat.VecDense, error) {

	// A_d = exp(A*dt)
	ad, err := discretize(d.A, dt)
	if err != nil {
		return nil, errors.New("discretization of A failed")
	}
	//fmt.Println("A_d=", ad)

	// B_d = Int_0^T exp(A*dt) * B dt
	bd, _ := integrate(ad, d.A, d.B, dt)
	//fmt.Println("B_d=", bd)

	// x(k+1) = A_d * x + B_d * u
	var xk1, adx, bdu mat.VecDense
	adx.MulVec(ad, x)
	bdu.MulVec(bd, u)
	xk1.AddVec(&adx, &bdu)

	return &xk1, nil
}

//Response returns the output vector y(t) = C * x(t) + D * u(t)
func (d *Discrete) Response(x *mat.VecDense, u *mat.VecDense) *mat.VecDense {

	// cx = C * x
	var cx mat.VecDense
	cx.MulVec(d.C, x)

	// du = D * u
	var du mat.VecDense
	du.MulVec(d.D, u)

	// y(t) = C * x(t) + D * u(t)
	var yk mat.VecDense
	yk.AddVec(&cx, &du)

	return &yk
}
