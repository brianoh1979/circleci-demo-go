package bigdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoComputation(t *testing.T) {
	result := DoComputation()
	assert.True(t, result)
}

func TestDoComplexComputation(t *testing.T) {
	result := DoComplexComputation()
	assert.False(t, result)
}

func TestDoComputation2(t *testing.T) {
        result := DoComputation()
        assert.True(t, result)
}

func TestDoComplexComputation2(t *testing.T) {
        result := DoComplexComputation()
        assert.True(t, result)
}
