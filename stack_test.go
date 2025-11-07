package kutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortStack(t *testing.T) {
	shortStack := ShortStack()
	t.Log(string(shortStack))
	assert.Equal(t, "testing.", string(shortStack[:8]))
	assert.EqualValues(t, '\n', shortStack[len(shortStack)-1])
}
