package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOwner(t *testing.T) {
	str := RandomString(6)

	require.NotEmpty(t, str)
	require.Len(t, str, 6)
}
