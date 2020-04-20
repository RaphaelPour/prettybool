package styles

import (
	"fmt"
)

type style struct {
	ok, notOk string
}

var (
	styles = map[string]style{
		"ok":       {ok: "ok", notOk: "not ok"},
		"yes":      {ok: "yes", notOk: "no"},
		"true":     {ok: "true", notOk: "false"},
		"0":        {ok: "0", notOk: "1"},
		"passed":   {ok: "passed", notOk: "failed"},
		"positive": {ok: "positive", notOk: "negative"},
		"good":     {ok: "good", notOk: "bad"},
		"check":    {ok: "\u2705", notOk: "\u274c"},
		"enabled":  {ok: "enabled", notOk: "disabled"},
		"on":       {ok: "on", notOk: "off"},
		"active":   {ok: "active", notOk: "inactive"},
		"success":  {ok: "success", notOk: "failure"},
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

func ListStyles() {
	fmt.Println("Pretty Boolean styles:")

	for name, s := range styles {
		fmt.Printf("%s: %s, %s\n", name, s.ok, s.notOk)
	}
}
