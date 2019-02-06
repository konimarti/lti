# State-Space Representation of Linear, Time-Invariant (LTI) Systems in Golang

[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/konimarti/lti/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/konimarti/observer?status.svg)](https://godoc.org/github.com/konimarti/lti)
[![goreportcard](https://goreportcard.com/badge/github.com/konimarti/observer)](https://goreportcard.com/report/github.com/konimarti/lti)

```go get github.com/konimarti/lti```

## Usage
```go
	// define system type (state-space model)
	system, err := lti.NewDiscrete(
		...
	)

	// define initial state (x) and control (u) vectors
	...

	// propagate LTI system for time step dt
	nextState, err := system.Propagate(dt, x, u)

	// get derivative vector for new state
	fmt.Println(system.Derivative(nextState, u))
	
	// get output vector for new state
	fmt.Println(system.Response(nextState, u))

}
```

## State-space models

The following system types are implemented:
* ```lti.Discrete{}```: Discrete, time-invariant

### Example

See example [here](example/lti.go).

## Credits

This software package has been developed for and is in production at [Kalkfabrik Netstal](http://www.kfn.ch/en).
