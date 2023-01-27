package math

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSumPositive(t *testing.T) {
	total := Sum(1, 1)
	assert.Equal(t, total, 2)
}
func TestSumNegative(t *testing.T) {
	total := Sum(1, -10)
	assert.Equal(t, total, -9)
}

func TestSumPositive2(t *testing.T) {
        total := Sum(1, 2)
        assert.Equal(t, total, 3) 
}
func TestSumNegative2(t *testing.T) {
        total := Sum(1, -11)
        assert.Equal(t, total, -10)
}

func TestSumLarge(t *testing.T) {
	time.Sleep(5 * time.Second)
	total := Sum(10000, 10000)
	assert.Equal(t, total, 20000)
}
