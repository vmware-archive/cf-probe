package helpers

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/cloudfoundry/gunk/command_runner"
	"github.com/pivotal-cf-experimental/cf-test-helpers/generator"
)

func NewBigApp(runner command_runner.CommandRunner, appPath string, sizeInMegabytes int) (*BigApp, error) {
	tempDirName, err := ioutil.TempDir("", "big-app")
	if err != nil {
		return nil, err
	}

	cmd := &exec.Cmd{Path: "cp", Args: []string{"-r", appPath + "/", tempDirName}}

	err = runner.Run(cmd)
	if err != nil {
		return nil, err
	}

	ddOutputArg := "of=" + tempDirName + "/payload"
	ddCountArg := fmt.Sprintf("count=%d", sizeInMegabytes)

	cmd = &exec.Cmd{Path: "dd", Args: []string{"if=/dev/urandom", ddOutputArg, ddCountArg, "bs=1048576"}}

	err = runner.Run(cmd)
	if err != nil {
		return nil, err
	}

	return &BigApp{
		Location: tempDirName,
		runner:   runner,
		Name:     "big-app-" + generator.RandomName(),
	}, nil
}

func (app *BigApp) Push() error {
	cmd := &exec.Cmd{Path: "gcf", Args: []string{"push", app.Name, "-p", app.Location}}
	return app.runner.Run(cmd)
}

type BigApp struct {
	Location string
	runner   command_runner.CommandRunner
	Name     string
}
