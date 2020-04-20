package styles

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInvalidStyle(t *testing.T) {

	pb := GetPrettyBool(true, "")
	require.Equal(t, pb, "")
}

func TestValidStyleWithTrue(t *testing.T) {
	pb := GetPrettyBool(true, "ok")
	require.Equal(t, pb, "ok")
}

func TestValidStyleWithFalse(t *testing.T) {
	pb := GetPrettyBool(false, "ok")
	require.Equal(t, pb, "not ok")
}
