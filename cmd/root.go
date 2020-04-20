package cmd // import "github.com/3vilcookie/prettybool/cmd"

import (
	"flag"
	"fmt"

	"github.com/3vilcookie/prettybool"
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
		styles.ListStyles()
		return
	}

	fmt.Println(styles.GetPrettyBool(*value, *style))
}
