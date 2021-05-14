package foobar

import (
	"github.com/stretchr/testify/require"
	"testing"
)
func TestCount(t *testing.T) {
	s := "qwerasdfe"
	require.Equal(t, Count(s, 'e'), 2, "counting 'e' in "+s)
	require.Equal(t, Count(s, 'x'), 0, "counting 'x' in "+s)
	require.Equal(t, Count(s, 'f'), 0, "counting 'f' in "+s)
}
