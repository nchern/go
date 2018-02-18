package limiters

import (
	"errors"
	"time"
)

var (
	defaultTimer = &defaultTimerImpl{}

	// ErrLimitExceeded signals that the counter limit has hit
	ErrLimitExceeded = errors.New("Limit exceeded")
)

// Counter abstracts counter interface
type Counter interface {
	IncrBy(int64) error
	Total() (int64, error)
}

// timer just abstracts time provider, this helps during testing
type timer interface {
	UnixNano() int64
}

type defaultTimerImpl struct{}

func (t *defaultTimerImpl) UnixNano() int64 { return time.Now().UnixNano() }

// CheckAndIncrement is a simple strategy to check a counter value for the limit and
// increment it if the counter value is still under the limit.
// NOTE: This strategy is not precise due to race condition but this is sufficient in most cases.
func CheckAndIncrement(counter Counter, incrVal, maxVal int64) error {
	val, err := counter.Total()
	if err != nil {
		return err
	}
	if val+incrVal > maxVal {
		return ErrLimitExceeded
	}
	return counter.IncrBy(incrVal)
}
