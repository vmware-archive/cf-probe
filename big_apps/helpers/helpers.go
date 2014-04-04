package helpers

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/cloudfoundry/gunk/command_runner"
)

func PushApp(runner command_runner.CommandRunner, name string, path string) error {
	cmd := &exec.Cmd{Path: "gcf", Args: []string{"push", name, "-p", path}}
	return runner.Run(cmd)
}

func MakeBigApp(runner command_runner.CommandRunner, appPath string, sizeInMegabytes int) (string, error) {
	tempDirName, err := ioutil.TempDir("", "big-app")
	if err != nil {
		return "", err
	}

	cmd := exec.Command("cp", "-r", appPath+"/", tempDirName)

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	ddOutputArg := "of=" + tempDirName + "/payload"
	ddCountArg := fmt.Sprintf("count=%d", sizeInMegabytes)

	cmd = exec.Command("dd", "if=/dev/urandom", ddOutputArg, ddCountArg, "bs=1048576")

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return tempDirName, nil
}
