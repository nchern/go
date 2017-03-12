package limiters

import (
	"testing"
	"time"

	"gopkg.in/stretchr/testify.v1/assert"
)

type redisConnMock struct {
	orderedKeys []string
	cache       map[string]int64
}

func newRedis() *redisConnMock {
	return &redisConnMock{
		cache: make(map[string]int64),
	}
}

func (r *redisConnMock) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if commandName == "GET" {
		return r.cache[args[0].(string)], nil
	}
	return 0, nil
}

func (r *redisConnMock) KeysCount() int {
	return len(r.cache)
}

func (r *redisConnMock) Send(commandName string, args ...interface{}) error {
	if commandName == "INCRBY" {
		k, val := args[0], args[1]
		key := k.(string)
		r.cache[key] += val.(int64)
		r.orderedKeys = append(r.orderedKeys, key)
	}
	return nil
}

func (c *redisConnMock) Close() error {
	panic("not implemented")
}

func (c *redisConnMock) Err() error {
	panic("not implemented")
}

func (c *redisConnMock) Flush() error {
	panic("not implemented")
}

func (c *redisConnMock) Receive() (reply interface{}, err error) {
	panic("not implemented")
}

func TestRedisCounter(t *testing.T) {
	conn := newRedis()

	limiter := NewRateRedisCounter(conn, "tst", time.Minute, 10*time.Second)

	tm := &testTimer{val: 1 * time.Second}
	limiter.timer = tm

	assert.Equal(t, 6, limiter.bucketsCount)

	limiter.IncrBy(6)
	total, _ := limiter.Total()
	assert.Equal(t, int64(6), total)
	// assert key structure
	assert.Equal(t, "cnt:tst:0", conn.orderedKeys[0])

	tm.val += 4 * time.Second // 5th second
	limiter.IncrBy(2)
	total, _ = limiter.Total()

	assert.Equal(t, int64(8), total)
	assert.Equal(t, 1, conn.KeysCount())

	tm.val += 6 * time.Second // 11th second - next bucket
	limiter.IncrBy(3)
	total, _ = limiter.Total()

	assert.Equal(t, int64(11), total)
	assert.Equal(t, 2, conn.KeysCount())

	tm.val += 21 * time.Second // 32d second
	limiter.IncrBy(1)
	total, _ = limiter.Total()

	assert.Equal(t, int64(12), total)
	assert.Equal(t, 3, conn.KeysCount())

	tm.val += 33 * time.Second // 64th second: one minute round
	limiter.IncrBy(5)
	total, _ = limiter.Total()

	assert.Equal(t, int64(9), total)
}
