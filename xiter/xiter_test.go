package xiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMisc(t *testing.T) {
	assert.Equal(t, 55, Reduce(
		func(acc int, n int) int { return acc + n },
		10,
		Take(Range(0), 10),
	))
}
