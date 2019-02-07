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
	A *mat.Dense
	B *mat.Dense
	C *mat.Dense
	D *mat.Dense
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

	// A * x(t)
	var ax mat.VecDense
	ax.MulVec(s.A, x)

	// B * u(t)
	var bx mat.VecDense
	bx.MulVec(s.B, u)

	// xderiv = A x(t) + B u(t)
	var xderiv mat.VecDense
	xderiv.AddVec(&ax, &bx)

	return &xderiv
}

//Response returns the output vector y(t) = C * x(t) + D * u(t)
func (s *System) Response(x *mat.VecDense, u *mat.VecDense) *mat.VecDense {

	// cx = C * x
	var cx mat.VecDense
	cx.MulVec(s.C, x)

	// du = D * u
	var du mat.VecDense
	du.MulVec(s.D, u)

	// y(t) = C * x(t) + D * u(t)
	var yk mat.VecDense
	yk.AddVec(&cx, &du)

	return &yk
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

	// controllability matrix
	n, _ := s.B.Dims()

	var c, ab mat.Dense
	c.Clone(s.B)
	ab.Clone(s.B)

	// create augmented matrix
	for i := 0; i < n-1; i++ {
		ab.Mul(s.A, &ab)
		var tmp mat.Dense
		tmp.Augment(&c, &ab)
		c.Clone(&tmp)
	}
	//fmt.Println(c)

	// calculate rank
	rank, err := rank(&c)
	if err != nil {
		return false, err
	}
	//fmt.Println("rank(C)=", rank)

	// check
	if rank < n {
		return false, nil
	}
	return true, nil
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

	// observability matrix S
	_, n := s.C.Dims()

	var sb, ca mat.Dense
	sb.Clone(s.C)
	ca.Clone(s.C)

	// create stacked matrix
	for i := 0; i < n-1; i++ {
		ca.Mul(&ca, s.A)
		var tmp mat.Dense
		tmp.Stack(&sb, &ca)
		sb.Clone(&tmp)
	}
	//fmt.Println("S=", s)

	// calculate rank
	rank, err := rank(&sb)
	if err != nil {
		return false, err
	}
	//fmt.Println("rank(S)=", rank)

	// check
	if rank < n {
		return false, nil
	}
	return true, nil
}
