package main

import (
	"flag"
	"fmt"

	"github.com/3vilcookie/prettybool/pkg/styles"
)

var (
	buildDate    string
	buildVersion string
	version      = flag.Bool("version", false, "Print build information")
	list         = flag.Bool("list", false, "Lists all semantic boolean style")
	style        = flag.String("style", "true", "Sets the output style. The name of the style is equal to the representation of the 'true' value.")
	value        = flag.Bool("value", true, "Boolean value which should be printed with another style")
)

func main() {

	flag.Parse()
	if *version {
		fmt.Printf("BuildDate: %s\n", buildDate)
		fmt.Printf("BuildVersion: %s\n", buildVersion)
		return
	}

	if *list {
		fmt.Println("Pretty Boolean styles:")
		fmt.Println("     yes: yes, no")
		fmt.Println("      ok: ok, not ok")
		fmt.Println("    okay: okay, not okay")
		fmt.Println("    true: true, false")
		fmt.Println("       t: t, f")
		fmt.Println("       0: 0, 1")
		fmt.Println("  passed: passed, failed")
		fmt.Println("positive: positive, negative")
		fmt.Println("    good: good, bad")
		fmt.Println("   check: \u2705, \u274c")
		return
	}

	if *style == "" {
		fmt.Printf("Style is missing")
		return
	}

	b, err := styles.GetPrettyBool(*value, *style)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(b)
}
