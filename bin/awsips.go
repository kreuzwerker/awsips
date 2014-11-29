package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	c "github.com/kreuzwerker/awsips"
)

var build string

func main() {

	var (
		help           = flag.Bool("h", false, "display usage")
		mode           = flag.String("m", "", "mode - either 'regions', 'services', 'list-files' or 'list-tree'")
		regionFilters  = flag.String("rf", "", "filter by region or multiple, comma-seperated regions")
		serviceFilters = flag.String("sf", "", "filter by service or multiple, comma-seperated services")
	)

	flag.Parse()

	if *help {
		printUsage()
	}

	out, err := c.Client()

	if err != nil {
		fail(err)
	}

	switch *mode {

	case "regions":
		fmt.Println(format(out.Regions()))

	case "services":
		fmt.Println(format(out.Services()))

	case "list-files":

		for service, regionTree := range out.Filter(filter(regionFilters), filter(serviceFilters)) {

			for region, ips := range regionTree {

				file := fmt.Sprintf("%s-%s.json", service, region)

				if err := ioutil.WriteFile(file, []byte(format(ips)), 0644); err != nil {
					fail(err)
				}

			}

		}

	case "list-tree":
		fmt.Println(format(out.Filter(filter(regionFilters), filter(serviceFilters))))

	default:
		printUsage()
	}

}

func fail(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func filter(f *string) []string {

	if *f == "" {
		return []string{}
	} else {
		return strings.Split(*f, ",")
	}

}

func format(v interface{}) string {

	out, err := json.MarshalIndent(v, "", "    ")

	if err != nil {
		fail(err)
	}

	return string(out)

}

func printUsage() {

	fmt.Fprintf(os.Stderr, "Usage of %s (%s):\n", os.Args[0], build)
	flag.PrintDefaults()

	os.Exit(1)

}
