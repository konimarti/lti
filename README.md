# Linear and time-invariant systems in Golang 

[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/konimarti/lti/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/konimarti/observer?status.svg)](https://godoc.org/github.com/konimarti/lti)
[![goreportcard](https://goreportcard.com/badge/github.com/konimarti/observer)](https://goreportcard.com/report/github.com/konimarti/lti)

```go get github.com/konimarti/lti```

* State-space representation and estimation of linear, time-invariant systems for control theory in Golang

	```math
	 x'(t) = A * x(t) + B * u(t)
	```
	 and
	```math
	 y(t)  = C * x(t) + D * u(t)
	```
* Can be used as an input for a [Kalman filter](http://github.com/konimarti/kalman). 

## Usage
```go
	// define time-continuous linear system
	system, err := lti.NewSystem(
		...
	)

	// check system properties
	fmt.Println("Observable=", system.MustObservable())
	fmt.Println("Controllable=", system.MustControllable())

	// define initial state (x) and control (u) vectors
	...

	// get derivative vector for new state
	fmt.Println(system.Derivative(x, u))

	// get output vector for new state
	fmt.Println(system.Response(x, u))

	// discretize LTI system and propagate state by time step dt
	discrete, err := system.Discretize(dt)

	fmt.Println("x(k+1)=", discrete.Predict(x, u))
}
```

### Example

See example [here](example/lti.go).

## More information

For additional materials on state-space models for control theory, check out the following links:
* A practical guide to state-space control [here](https://github.com/calcmogul/state-space-guide)
* State-space model impelmentation for Arduinos [here](https://github.com/tomstewart89/StateSpaceControl)

## Credits

This software package has been developed for and is in production at [Kalkfabrik Netstal](http://www.kfn.ch/en).
