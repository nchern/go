package main

type asyncProcessorFixed struct {
	processorsCount int
	q               chan<- *request
}

// NewAsyncProcessorFixed creates AsyncProcessor that has fixed number of goroutines to process calls
func NewAsyncProcessorFixed(n int) AsyncProcessor {
	return &asyncProcessorFixed{
		processorsCount: n,
		q:               start(n),
	}
}

func (p *asyncProcessorFixed) Call(action Action) AsyncResult {
	req := newRequest(action)
	p.q <- req
	return req
}

func (p *asyncProcessorFixed) CallMany(actions ...Action) AsyncResults {
	reqs := make([]*request, len(actions))
	for i, action := range actions {
		req := newRequest(action)
		reqs[i] = req
		p.q <- req
	}
	return requests(reqs)
}

func main() {

	//var result []byte
	processor := NewAsyncProcessorFixed(10)
	processor.Call(func() error {
		// http.Get
		return nil
	}).Wait()
}
