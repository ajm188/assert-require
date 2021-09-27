package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssert(t *testing.T) {
	assert.Equal(t, 1, 2)
	assert.Equal(t, 2, 3)
	assert.Equal(t, 4, 4)
}

func TestRequire(t *testing.T) {
	require.Equal(t, 1, 2)
	require.Equal(t, 2, 3)
	require.Equal(t, 4, 4)
}
