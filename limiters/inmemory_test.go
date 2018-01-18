package limiters

import (
	"testing"
	"time"

	"gopkg.in/stretchr/testify.v1/assert"
)

type testTimer struct {
	val time.Duration
}

func (c *testTimer) UnixNano() int64 {
	return int64(c.val)
}

func Test(t *testing.T) {

	empty := bucket{value: 0, timeShard: 0}

	limiter := NewRateBucketCounter(time.Minute, 10*time.Second)

	tm := &testTimer{val: 10 * time.Second}
	limiter.timer = tm

	assert.Equal(t, 6, len(limiter.buckets))

	limiter.IncrBy(6)
	total, _ := limiter.Total()
	assert.Equal(t, int64(6), total)
	assert.Equal(t,
		[]bucket{empty, bucket{value: 6, timeShard: 1}, empty, empty, empty, empty},
		limiter.buckets)

	tm.val += 4 * time.Second // 5th second
	limiter.IncrBy(2)
	total, _ = limiter.Total()
	assert.Equal(t, int64(8), total)
	assert.Equal(t,
		[]bucket{empty, bucket{value: 8, timeShard: 1}, empty, empty, empty, empty},
		limiter.buckets)

	tm.val += 6 * time.Second // 11th second
	limiter.IncrBy(3)
	total, _ = limiter.Total()
	assert.Equal(t, int64(11), total)
	assert.Equal(t,
		[]bucket{empty, bucket{value: 8, timeShard: 1}, bucket{3, 2}, empty, empty, empty},
		limiter.buckets)

	tm.val += 21 * time.Second // 32d second
	limiter.IncrBy(1)
	total, _ = limiter.Total()
	assert.Equal(t, int64(12), total)
	assert.Equal(t,
		[]bucket{empty, bucket{value: 8, timeShard: 1}, bucket{3, 2}, empty, bucket{1, 4}, empty},
		limiter.buckets)

	tm.val += 33 * time.Second // 64th second: one minute round
	limiter.IncrBy(5)
	total, _ = limiter.Total()
	assert.Equal(t, int64(9), total)
	assert.Equal(t,
		[]bucket{empty, bucket{value: 5, timeShard: 7}, bucket{3, 2}, empty, bucket{1, 4}, empty},
		limiter.buckets)

	tm.val += 60 * time.Second // 124th second: two minutes passed
	limiter.IncrBy(2)
	total, _ = limiter.Total()
	assert.Equal(t, int64(2), total)
	assert.Equal(t,
		[]bucket{empty, bucket{value: 2, timeShard: 13}, bucket{3, 2}, empty, bucket{1, 4}, empty},
		limiter.buckets)

	tm.val += 80 * time.Second // 214th second: 3 minutes passed
	limiter.IncrBy(5)
	total, _ = limiter.Total()
	assert.Equal(t, int64(5), total)

	tm.val += 12 * time.Second // 226th second: 3 minutes passed
	limiter.IncrBy(3)
	total, _ = limiter.Total()
	assert.Equal(t, int64(8), total)
}
