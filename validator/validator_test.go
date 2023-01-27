package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	result := Validate("my string")
	assert.True(t, result)
}

func TestValidate2(t *testing.T) {
        result := Validate("my string")
        assert.True(t, result)
}

func TestValidate3(t *testing.T) {
        result := Validate("my string")
        assert.True(t, result)
}

func TestValidate4(t *testing.T) {
        result := Validate("my string")
        assert.True(t, result)
}

func TestValidate5(t *testing.T) {
        result := Validate("my string")
        assert.True(t, result)
}
