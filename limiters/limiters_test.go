package limiters

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCounter struct {
	value int64
	err   error
}

func (c *mockCounter) IncrBy(val int64) error {
	c.value += val
	return c.err
}

func (c *mockCounter) Total() (int64, error) {
	return c.value, c.err
}

func TestCheckAndIncrementWithCounterErrors(t *testing.T) {
	counter := &mockCounter{}
	counter.err = errors.New("Some error")

	assert.Equal(t, counter.err, CheckAndIncrement(counter, 10, 15))
	assert.Equal(t, int64(0), counter.value)
}

func TestCheckAndIncrement(t *testing.T) {
	counter := &mockCounter{}

	assert.Nil(t, CheckAndIncrement(counter, 10, 15))
	assert.Equal(t, int64(10), counter.value)

	// hit the limit
	assert.Equal(t, ErrLimitExceeded, CheckAndIncrement(counter, 10, 15))
	assert.Equal(t, int64(10), counter.value)
}
