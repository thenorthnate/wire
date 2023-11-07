package wire

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	definition, err := parse("examples/forest.toml")
	require.NoError(t, err)
	fmt.Println(definition)
}
