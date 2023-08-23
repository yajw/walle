package lpr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func dateOf(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
}

func TestLPR1(t *testing.T) {
	assert.Equal(t, Get5Y(dateOf(2018, 8, 2)).String(), "0.0485")
	assert.Equal(t, Get5Y(dateOf(2023, 8, 2)).String(), "0.042")
	assert.Equal(t, Get5Y(dateOf(2020, 8, 2)).String(), "0.0465")
	assert.Equal(t, Get1Y(dateOf(2018, 8, 2)).String(), "0.0425")
}
