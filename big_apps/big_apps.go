package main

import (
	"flag"
	"fmt"
	// "os"

	"github.com/cloudfoundry/gunk/command_runner"
	"github.com/pivotal-cf-experimental/cf-probe/big_apps/helpers"
)

var assetPath = flag.String("app-path", "../assets/big-app-base", "path to app")

func main() {
	flag.Parse()

	runner := command_runner.New(true)

	app, err := helpers.NewBigApp(runner, *assetPath, 5)
	if err != nil {
		fmt.Printf("Creating app failed")
		return
	}

	err = app.Push()
	if err != nil {
		fmt.Printf("App push failed")
		return
	}

	return
}
