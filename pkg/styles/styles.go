package styles

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type style struct {
	ok, notOk string
}

var (
	styles = map[string]style{
		"ok":  {ok: "ok", notOk: "not ok"},
		"yes": {ok: "yes", notOk: "no"},
	}
)

func GetPrettyBool(value bool, styleName string) (string, error) {

	s, ok := styles[styleName]

	if !ok {
		return "", fmt.Errorf("Unknown style %s", styleName)
	}

	if value {
		return s.ok, nil
	}

	return s.notOk, nil
}
