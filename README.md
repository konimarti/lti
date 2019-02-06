# Linear and time-invariant (LTI) systems for control theory in Golang

[![License](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/konimarti/lti/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/konimarti/observer?status.svg)](https://godoc.org/github.com/konimarti/lti)
[![goreportcard](https://goreportcard.com/badge/github.com/konimarti/observer)](https://goreportcard.com/report/github.com/konimarti/lti)

```go get github.com/konimarti/lti```

* State-space representation and estimation of linear, time-invariant (LTI) systems for control theory in Golang

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

### Example

See example [here](example/lti.go).

## State-space models

The following system types are implemented:
* ```lti.Discrete{}```: Discrete, time-invariant

### Understanding state-space models for control theory

For additional materials on state-space models check out the following links:
* A practical to state-space control [here](https://github.com/calcmogul/state-space-guide)
* State-space model impelmentation for Arduinos [here](https://github.com/tomstewart89/StateSpaceControl)

## Credits

This software package has been developed for and is in production at [Kalkfabrik Netstal](http://www.kfn.ch/en).
