package main

import (
	"fmt"
	"time"
)

var (
	// ErrCallTimedOut indicates that timeout occured for the call
	ErrCallTimedOut = fmt.Errorf("Timeout occured for the call")
)

// Action represents a basic function type
type Action func() error

// AsyncResult abstracts a result of async operation
type AsyncResult interface {
	Wait() error

	WaitWithTimeout(timeout time.Duration) error
}

type AsyncResults interface {
	WaitAll() []error

	WaitAllWithTimeout(timeout time.Duration) []error
}

// AsyncProcessor abstracts the asyncronious function caller
type AsyncProcessor interface {
	Call(action Action) AsyncResult

	CallMany(action ...Action) AsyncResults
}

// CallWithTimeout is a convenient helper to call an arbitrary function
// implementing the following pattern: https://gobyexample.com/timeouts
func CallWithTimeout(action Action, timeout time.Duration) error {
	c := make(chan error)
	go func() {
		c <- action()
	}()

	select {
	case res := <-c:
		return res
	case <-time.After(timeout):
		return ErrCallTimedOut
	}
	panic("Should never happen")
}

type requests []*request

func (r requests) WaitAll() []error {
	res := make([]error, len(r))
	for i := range r {
		res[i] = r[i].Wait()
	}
	return res
}

func (r requests) WaitAllWithTimeout(timeout time.Duration) []error {
	// FIXME: incorrect implementation
	res := make([]error, len(r))
	for i := range r {
		res[i] = r[i].WaitWithTimeout(timeout)
		if res[i] == ErrCallTimedOut {
			for j := i + 1; j < len(r); j++ {
				res[j] = ErrCallTimedOut
			}
			return res
		}
	}
	return res
}

type request struct {
	action Action

	response chan error

	stop bool
}

func newRequest(action Action) *request {
	return &request{
		action:   action,
		response: make(chan error, 1),
	}
}

func (r *request) Wait() error {
	return <-r.response
}

func (r *request) WaitWithTimeout(timeout time.Duration) error {
	select {
	case res := <-r.response:
		return res
	case <-time.After(timeout):
		return ErrCallTimedOut
	}
	panic("Should never happen")
}

func process(queue <-chan *request) {
	for req := range queue {
		if req.stop {
			return
		}
		req.response <- req.action()
	}
}

func start(n int) chan *request {
	q := make(chan *request, 2)
	for i := 0; i < n; i++ {
		go process(q)
	}
	return q
}
