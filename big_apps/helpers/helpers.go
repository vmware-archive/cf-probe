package helpers

import (
	"github.com/cloudfoundry/gunk/command_runner"
)

type BinarySearchTester interface {
	Test(value int) bool
}

func BinarySearch(tester BinarySearchTester, low, high, tolerance int) int {
	for (high - low) > tolerance {
		testValue := (high + low) / 2
		result := tester.Test(testValue)
		if result {
			low = testValue
		} else {
			high = testValue
		}
	}

	return low
}

type AppSizeBinarySearchTester struct {
	runner  command_runner.CommandRunner
	appPath string
}

func (tester *AppSizeBinarySearchTester) Test(value int) bool {
	app, err := NewBigApp(tester.runner, tester.appPath, value)
	if err != nil {
		return false
	}

	// defer app.Cleanup()
	err = app.Push()
	return err == nil
}

func NewAppSizeBinarySearchTester(runner command_runner.CommandRunner, appPath string) *AppSizeBinarySearchTester {
	return &AppSizeBinarySearchTester{
		runner:  runner,
		appPath: appPath,
	}
}
