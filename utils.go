package lti

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

//Discretize
// A_discretized = exp(A * t)
func discretize(m *mat.Dense, t float64) (*mat.Dense, error) {
	// m_d = exp(m * t)

	// check if matrix m is square
	if r, c := m.Dims(); r != c {
		return m, errors.New("Discretize: matrix is not square")
	}

	// tmp = m * t
	var tmp mat.Dense
	tmp.Scale(t, m)

	// exp( tmp )
	var md mat.Dense
	md.Exp(&tmp)

	return &md, nil
}

// Integrate
// B_discretized = Int_0^T exp(A t) B dt
// with exp(A t) = Sum_k (A t)^k / k!
// Source: https://math.stackexchange.com/questions/658276/integral-of-matrix-exponential
func integrate(a *mat.Dense, b *mat.Dense, t float64) (*mat.Dense, error) {

	// B_d = Int_0^T exp(At) * B dt
	// B_d = (Int_0^T exp(At) dt) * B
	// B_d = ( T [ Sum_n (AT)^n-1 ] ) * B

	// (At)
	var at mat.Dense
	at.Scale(t, a)

	// Sum_n (AT)^n-1 / n!
	var x, tmp mat.Dense
	x.Clone(&at)
	x.Zero()
	fac := 1.0
	for n := 1; n < 10; n++ {
		// (AT)^n-1
		tmp.Pow(&at, n-1)
		// n!
		fac = fac * float64(n)
		tmp.Scale(1.0/fac, &tmp)
		//fmt.Println("n=", n, "fac=", fac, "tmp=", tmp)
		x.Add(&tmp, &x)
	}
	//fmt.Println("at=", at)
	//fmt.Println("x=", x)

	// Int * B
	var bd mat.Dense
	bd.Mul(&x, b)

	bd.Scale(t, &bd)

	return &bd, nil
}

// rank calculates rank of matrix using singular value decomposition
func rank(a *mat.Dense) (int, error) {
	var svd mat.SVD
	ok := svd.Factorize(a, mat.SVDNone)
	if !ok {
		return 0, errors.New("rank: factorization failed")
	}
	rank := 0
	for _, value := range svd.Values(nil) {
		if value > 1e-8 {
			rank += 1
		}
	}
	return rank, nil
}
