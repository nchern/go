package limiters

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RateRedisCounter struct {
	timer timer

	period     time.Duration
	resolution time.Duration

	bucketsCount int

	prefix string
	conn   redis.Conn
}

// NewRateRedisCounter initializes new Redis-based sharded counter
// prefix: a custom string prefix that will be added to the each bucket key
// period: defines the maximum period for which we count, e.g. 5m, 1h
// resolution: defines the resolution of a counter, i.e. the minimum time period of a counter bucket.
// counter value is calculated as a sum of all bucket values within the last period
func NewRateRedisCounter(conn redis.Conn, prefix string, period time.Duration, resolution time.Duration) *RateRedisCounter {
	assertPeriodAndResolutionCorrect(period, resolution)

	return &RateRedisCounter{
		conn:         conn,
		period:       period,
		prefix:       prefix,
		resolution:   resolution,
		timer:        defaultTimer,
		bucketsCount: int(period / resolution),
	}
}

func (c *RateRedisCounter) IncrBy(val int64) error {
	curTimeShard := c.timer.UnixNano() / int64(c.resolution)
	key := c.makeKey(curTimeShard)

	c.conn.Send("MULTI")
	c.conn.Send("INCRBY", key, val)
	c.conn.Send("EXPIRE", key, int(c.resolution/time.Second)+2) // give 2 secs more than the resolution

	_, err := c.conn.Do("EXEC")
	return err
}

func (c *RateRedisCounter) makeKey(timeShard int64) string {
	return fmt.Sprintf("cnt:%s:%d", c.prefix, timeShard)
}

func (c *RateRedisCounter) Total() (int64, error) {
	var total int64

	curTimeShard := c.timer.UnixNano() / int64(c.resolution)

	for i := 0; i < c.bucketsCount; i++ {
		key := c.makeKey(curTimeShard - int64(i))
		val, err := redis.Int64(c.conn.Do("GET", key))
		if err != nil {
			return 0, err
		}
		total += val
	}

	return total, nil
}
