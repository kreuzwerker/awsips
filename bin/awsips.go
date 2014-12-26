package main

import (
	"flag"
	"fmt"
	"os"

	c "github.com/kreuzwerker/awsips"
)

var build string

func main() {

	var (
		help           = flag.Bool("h", false, "display usage")
		mode           = flag.String("m", "", "mode - either 'regions', 'services', 'list-files', 'list-tree' or 'list-stack'")
		ports          = flag.String("p", "", fmt.Sprintf("ports for stack mode (defaults to using %s - %s)", c.DefaultFromPortForStack, c.DefaultToPortForStack))
		regionFilters  = flag.String("rf", "", "filter by region or multiple, comma-seperated regions")
		serviceFilters = flag.String("sf", "", "filter by service or multiple, comma-seperated services")
	)

	flag.Parse()

	if *help {
		printUsage()
	}

	result, err := c.Client()

	if err != nil {
		fail(err)
	}

	out := &c.Output{result}

	switch *mode {

	case "regions":
		output(out.Regions())

	case "services":
		output(out.Services())

	case "list-files":
		if err := out.Files(regionFilters, serviceFilters); err != nil {
			fail(err)
		}

	case "list-tree":
		output(out.Tree(regionFilters, serviceFilters))

	case "list-stack":
		output(out.Stack(regionFilters, serviceFilters, ports))

	default:
		printUsage()
	}

}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func output(out string, err error) {

	if err != nil {
		fail(err)
	}

	fmt.Println(out)

}

func printUsage() {

	fmt.Fprintf(os.Stderr, "Usage of %s (%s):\n", os.Args[0], build)
	flag.PrintDefaults()

	os.Exit(0)

}
