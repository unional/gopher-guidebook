package syntax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorialFor(t *testing.T) {
	result := FactorialFor(4)

	assert.Equal(t, 24, result)
}

func TestFactorialWhile(t *testing.T) {
	result := FactorialWhile(4)

	assert.Equal(t, 24, result)
}

func TestRandomizeString(t *testing.T) {
	result := RandomizeString("info", 2)

	assert.Equal(t, "info", result)
}
