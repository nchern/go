package limiters

import (
	"fmt"
	"sync"
	"time"
)

func assertPeriodAndResolutionCorrect(period, resolution time.Duration) {
	if period < resolution {
		panic(fmt.Sprintf("period[%v] must be greater or equal to resolution[%v]", period, resolution))
	}
}

type RateBucketCounter struct {
	prevBucketNo int64
	buckets      []int64
	lock         sync.RWMutex

	period     int64
	resolution int64

	timer timer
}

// NewRateBucketCounter initializes new in-memory sharded counter
// period: defines the maximum period for which we count, e.g. 5m, 1h
// resolution: defines the resolution of a counter, i.e. the minimum time period of a counter bucket.
// counter value is calculates as a sum of all bucket values within the last period
func NewRateBucketCounter(period time.Duration, resoultion time.Duration) *RateBucketCounter {
	assertPeriodAndResolutionCorrect(period, resoultion)
	return &RateBucketCounter{
		timer:      defaultTimer,
		period:     int64(period),
		resolution: int64(resoultion),
		buckets:    make([]int64, int(period/resoultion)),
	}
}

func (c *RateBucketCounter) IncrBy(val int64) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	bucketNo := (c.timer.UnixNano() / c.resolution) % int64(len(c.buckets))
	if bucketNo != c.prevBucketNo {
		c.buckets[bucketNo] = 0
	}
	c.prevBucketNo = bucketNo

	c.buckets[bucketNo] += val

	return nil
}

func (c *RateBucketCounter) Total() (int64, error) {
	total := int64(0)

	c.lock.RLock()
	defer c.lock.RUnlock()

	for _, v := range c.buckets {
		total += v
	}
	return total, nil
}
