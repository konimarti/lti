package lti

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

//IdealDiscretization returns a discretized matrix Md = exp(A*t) * M.
//If M is nil, then it just returns exp(A*t).
func IdealDiscretization(A *mat.Dense, dt float64, M *mat.Dense) (*mat.Dense, error) {
	// A_d = exp(A*dt)
	d, err := discretize(A, dt)
	if err != nil {
		return nil, errors.New("discretization of A failed")
	}

	// M_d = exp(A*dt) * M
	var md mat.Dense
	if M != nil {
		md.Mul(d, M)
	} else {
		md.CloneFrom(d)
	}
	return &md, nil
}

//RealDiscretization returns a discretized matrix Md = Int_0^T exp(A*t) * M dt.
func RealDiscretization(A *mat.Dense, dt float64, M *mat.Dense) (*mat.Dense, error) {
	// M_d = Int_0^T exp(A*dt) * M dt
	md, err := integrate(A, M, dt)
	if err != nil {
		return nil, errors.New("discretization of M failed")
	}

	return md, nil
}
