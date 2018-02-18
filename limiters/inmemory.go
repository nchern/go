package limiters

import (
	"fmt"
	"sync"
	"time"
)

type bucket struct {
	value     int64
	timeShard int64
}

func assertPeriodAndResolutionCorrect(period, resolution time.Duration) {
	if period < resolution {
		panic(fmt.Sprintf("period[%v] must be greater or equal to resolution[%v]", period, resolution))
	}
}

// RateBucketCounter represents in-memory sharded counter
type RateBucketCounter struct {
	buckets []bucket

	lock sync.RWMutex

	period     int64
	resolution int64

	timer timer
}

// NewRateBucketCounter initializes new in-memory sharded counter
// period: defines the maximum period for which we count, e.g. 5m, 1h
// resolution: defines the resolution of a counter, i.e. the minimum time period of a counter bucket.
// counter value is calculated as a sum of all bucket values within the last period
func NewRateBucketCounter(period time.Duration, resoultion time.Duration) *RateBucketCounter {
	assertPeriodAndResolutionCorrect(period, resoultion)
	return &RateBucketCounter{
		timer:      defaultTimer,
		period:     int64(period),
		resolution: int64(resoultion),
		buckets:    make([]bucket, int(period/resoultion)),
	}
}

// IncrBy adds the given value to this counter
func (c *RateBucketCounter) IncrBy(val int64) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	curTimeShard := c.timer.UnixNano() / c.resolution
	bucketNo := curTimeShard % int64(len(c.buckets))

	if c.buckets[bucketNo].timeShard == 0 || curTimeShard-c.buckets[bucketNo].timeShard >= c.period/c.resolution {
		c.buckets[bucketNo].value = 0
		c.buckets[bucketNo].timeShard = curTimeShard
	}

	c.buckets[bucketNo].value += val

	return nil
}

// Total returns the total value of this counter
func (c *RateBucketCounter) Total() (int64, error) {
	total := int64(0)

	c.lock.RLock()
	defer c.lock.RUnlock()

	curTimeShard := c.timer.UnixNano() / c.resolution
	for _, b := range c.buckets {
		if curTimeShard-b.timeShard > c.period/c.resolution {
			continue
		}
		total += b.value
	}
	return total, nil
}
