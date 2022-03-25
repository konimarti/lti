package lti

import "gonum.org/v1/gonum/mat"

//Covariance contains a discretized matrix to predict the next covariance matrix
// according to p(k+1) = Md * p(k) * Md^T
type Covariance struct {
	Md *mat.Dense
}

//NewCovariance creates a new covariance struct which
//needs to be initialized with a discretizied matrix Md
func NewCovariance(md *mat.Dense) *Covariance {
	return &Covariance{Md: md}
}

//Predict propagates the covariance p(k) to p(k+1)
//according to p(k+1) = Md * p(k) * Md^T;
//additionally noise is added if not nil
func (c *Covariance) Predict(p *mat.Dense, noise *mat.Dense, pmt, mpmt *mat.Dense) *mat.Dense {
	// p(k+1) = m * p(k) * m^T + noise
	pmt.Mul(p, c.Md.T())
	mpmt.Mul(c.Md, pmt)

	if noise != nil {
		mpmt.Add(mpmt, noise)
	}

	return mpmt
}
