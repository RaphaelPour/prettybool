package styles

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInvalidStyle(t *testing.T) {

	_, err := GetPrettyBool(true, "")
	require.NotNil(t, err)
}

func TestValidStyle(t *testing.T) {
	_, err := GetPrettyBool(true, "ok")
	require.Nil(t, err)
}
