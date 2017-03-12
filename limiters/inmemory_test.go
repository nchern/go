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

//func int_(i int64, err erro

func Test(t *testing.T) {
	limiter := NewRateBucketCounter(time.Minute, 10*time.Second)

	tm := &testTimer{val: 1 * time.Second}
	limiter.timer = tm

	assert.Equal(t, 6, len(limiter.buckets))

	limiter.IncrBy(6)
	total, _ := limiter.Total()
	assert.Equal(t, int64(6), total)
	assert.Equal(t, int64(6), limiter.buckets[0])

	tm.val += 4 * time.Second // 5th second
	limiter.IncrBy(2)
	total, _ = limiter.Total()
	assert.Equal(t, int64(8), total)
	assert.Equal(t, int64(8), limiter.buckets[0])

	tm.val += 6 * time.Second // 11th second
	limiter.IncrBy(3)
	total, _ = limiter.Total()
	assert.Equal(t, int64(11), total)
	assert.Equal(t, int64(8), limiter.buckets[0])
	assert.Equal(t, int64(3), limiter.buckets[1])

	tm.val += 21 * time.Second // 32d second
	limiter.IncrBy(1)
	total, _ = limiter.Total()
	assert.Equal(t, int64(12), total)
	assert.Equal(t, int64(8), limiter.buckets[0])
	assert.Equal(t, int64(3), limiter.buckets[1])
	assert.Equal(t, int64(0), limiter.buckets[2])
	assert.Equal(t, int64(1), limiter.buckets[3])

	tm.val += 33 * time.Second // 64th second: one minute round
	limiter.IncrBy(5)
	total, _ = limiter.Total()
	assert.Equal(t, int64(9), total)
	assert.Equal(t, int64(5), limiter.buckets[0])
	assert.Equal(t, int64(3), limiter.buckets[1])
	assert.Equal(t, int64(0), limiter.buckets[2])
	assert.Equal(t, int64(1), limiter.buckets[3])
}
