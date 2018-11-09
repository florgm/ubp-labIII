package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameFileDoesntExist(t *testing.T) {
	lines, err := ReadFile("abcd")

	assert.NotNil(t, err)
	assert.Empty(t, lines)
}
