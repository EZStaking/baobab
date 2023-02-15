package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrependJunoIfUnique(t *testing.T) {
	require := require.New(t)
	tcs := []struct {
		in  []string
		out []string
	}{
		// Should prepend "ubaobab" to a slice of denoms, unless it is already present.
		{[]string{}, []string{"ubaobab"}},
		{[]string{"a"}, []string{"ubaobab", "a"}},
		{[]string{"x", "a", "heeeyyy"}, []string{"ubaobab", "x", "a", "heeeyyy"}},
		{[]string{"x", "a", "ubaobab"}, []string{"x", "a", "ubaobab"}},
	}
	for i, tc := range tcs {
		require.Equal(tc.out, prependJunoIfUnique(tc.in), i)
	}
}
