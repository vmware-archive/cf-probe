package main

import (
	"flag"
	"fmt"
	// "os"

	"github.com/cloudfoundry/gunk/command_runner"
	"github.com/pivotal-cf-experimental/cf-probe/big_apps/helpers"
)

var assetPath = flag.String("app-path", "../assets/big-app-base", "path to app")
var lowLimit = flag.Int("low", 128, "lowest size to test")
var highLimit = flag.Int("high", 1024, "highest size to test")
var tolerance = flag.Int("tolerance", 4, "acceptable uncertainty in result")

func main() {
	flag.Parse()

	runner := command_runner.New(true)
	tester := helpers.NewAppSizeBinarySearchTester(runner, *assetPath)

	value := helpers.BinarySearch(tester, *lowLimit, *highLimit, *tolerance)
	fmt.Printf("The biggest app I was able to push was %dM", value)
	return
}
