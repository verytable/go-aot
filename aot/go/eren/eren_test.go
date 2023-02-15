package eren_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.verytable.online/aot/aot/go/eren"
)

func TestAbilities(t *testing.T) {
	require.NotEmpty(t, eren.Abilities())
	require.Contains(t, eren.Abilities(), "Power of the Titans")
}
