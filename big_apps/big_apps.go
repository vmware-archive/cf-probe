package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cloudfoundry/gunk/command_runner"
	"github.com/pivotal-cf-experimental/cf-probe/big_apps/helpers"
)

var assetPath = flag.String("app-path", "../assets/big-app-base", "path to app")

func main() {
	flag.Parse()

	runner := command_runner.New(true)

	appPath, err := helpers.MakeBigApp(runner, *assetPath, 1)
	if err != nil {
		fmt.Printf("Making big app failed: %s", err.Error())
	}

	defer os.RemoveAll(appPath)

	err = helpers.PushApp(runner, "big-app", appPath)
	if err != nil {
		fmt.Printf("Push failed: %s", err.Error())
	}
}
