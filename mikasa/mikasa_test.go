package mikasa_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.verytable.online/aot/aot/go/mikasa"
)

func TestAbilities(t *testing.T) {
	require.NotEmpty(t, mikasa.Abilities())
	require.Contains(t, mikasa.Abilities(), "Strength")
}
