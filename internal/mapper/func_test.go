package mapper

import (
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
)

func TestFunc(t *testing.T) {
	require := require.New(t)

	addTwo := func(a int) int { return a + 2 }
	f, err := NewFunc(addTwo)
	require.NoError(err)
	result, err := f.Call(1)
	require.NoError(err)
	require.Equal(result, 3)
}

func TestFunc_hclog(t *testing.T) {
	require := require.New(t)

	factory := func(log hclog.Logger) int { return 42 }
	f, err := NewFunc(factory)
	require.NoError(err)
	result, err := f.Call(hclog.L())
	require.NoError(err)
	require.Equal(result, 42)
}