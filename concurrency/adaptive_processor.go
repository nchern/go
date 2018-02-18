package main

import (
	"sync/atomic"
	"time"
)

type asyncProcessorAdaptive struct {
	processorsCount int64
	q               chan *request

	callsInFlight int64
}

// NewAsyncProcessorFixed creates AsyncProcessor that has fixed number of goroutines to process calls
func NewAdaptiveAsyncProcessor(n int) AsyncProcessor {
	res := &asyncProcessorAdaptive{
		processorsCount: int64(n),
		q:               start(n),
	}
	go res.recycleIdleProcessors()
	return res
}

func (p *asyncProcessorAdaptive) recycleIdleProcessors() {
	for range time.Tick(10 * time.Second) {
		idleCount := atomic.LoadInt64(&p.processorsCount) - atomic.LoadInt64(&p.callsInFlight)

		// TODO: abstract
		for i := int64(0); i < idleCount-5; i++ {
			p.q <- &request{stop: true}
		}
	}
}

func (p *asyncProcessorAdaptive) bumpUpCallsAndMaybeAddProcessor() {
	callsInFlight := atomic.AddInt64(&p.callsInFlight, 1)
	// TODO: abstract
	if atomic.LoadInt64(&p.processorsCount) < callsInFlight+1 {
		go process(p.q)
		atomic.AddInt64(&p.processorsCount, 1)
	}
}

func (p *asyncProcessorAdaptive) Call(action Action) AsyncResult {
	p.bumpUpCallsAndMaybeAddProcessor()
	defer atomic.AddInt64(&p.callsInFlight, -1)

	req := newRequest(action)
	p.q <- req
	return req
}

func (p *asyncProcessorAdaptive) CallMany(actions ...Action) AsyncResults {
	panic("not implemented")
}
