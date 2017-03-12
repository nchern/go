package limiters

import (
	"errors"
	"time"
)

var (
	defaultTimer = &defaultTimerImpl{}
)

type Counter interface {
	IncrBy(int64) error
	Total() (int64, error)
}

type timer interface {
	UnixNano() int64
}

type defaultTimerImpl struct {
}

func (t *defaultTimerImpl) UnixNano() int64 { return time.Now().UnixNano() }

var ErrLimitExceeded = errors.New("Limit exceeded")

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
