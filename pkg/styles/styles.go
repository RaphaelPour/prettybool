package styles

import (
	"fmt"
)

var (
	styles = map[string]map[bool]string{
		"ok":       {true: "ok", false: "not ok"},
		"yes":      {true: "yes", false: "no"},
		"true":     {true: "true", false: "false"},
		"0":        {true: "0", false: "1"},
		"passed":   {true: "passed", false: "failed"},
		"positive": {true: "positive", false: "negative"},
		"good":     {true: "good", false: "bad"},
		"check":    {true: "\u2705", false: "\u274c"},
		"enabled":  {true: "enabled", false: "disabled"},
		"on":       {true: "on", false: "off"},
		"active":   {true: "active", false: "inactive"},
		"success":  {true: "success", false: "failure"},
	}
)

func GetPrettyBool(value bool, styleName string) string {

	s, ok := styles[styleName]

	if !ok {
		return ""
	}

	return s[value]
}

func ListStyles() {
	fmt.Println("Pretty Boolean styles:")

	for name, s := range styles {
		fmt.Printf("%s: %s, %s\n", name, s[true], s[false])
	}
}
