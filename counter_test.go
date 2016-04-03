package gostats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter(t *testing.T) {
	v := perSecondCounter("counterTest", 10)
	assert.Equal(t, float64(0), v, "first counter should be zero")
}
