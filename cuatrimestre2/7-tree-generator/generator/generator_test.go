package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTreeOk(t *testing.T) {
	_, error := GenerateTree("../input.txt")

	assert.Nil(t, error)
}
