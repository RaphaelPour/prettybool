package main

import "fmt"

var (
	buildDate    string
	buildVersion string
	version      = flag.Bool("version", false, "Print build information")
	list         = flag.Bool("list", false, "Lists all semantic boolean style")
	style        = flag.String("style", "true", "Sets the output style. The name of the style is equal to the representation of the 'true' value.")
	value        = flag.Bool("value", true, "Boolean value which should be printed with another style")
)

func main() {

	if *version {
		fmt.Printf("BuildDate: %s\n", buildDate)
		fmt.Printf("BuildVersion: %s\n", buildVersion)
		return
	}

	if *list {
		fmt.Println("      ok: ok, not ok")
		fmt.Println("    okay: okay, not okay")
		fmt.Println("    true: true, false")
		fmt.Println("  passed: passed, failed")
		fmt.Println("     yes: yes, no")
		fmt.Println("       t: t, f")
		fmt.Println("       0: 0, 1")
		fmt.Println("positive: positive, negative")
		fmt.Println("    good: good, bad")
		return
	}
}
