package helpers_test

import (
	"errors"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/gunk/command_runner/fake_command_runner"

	. "github.com/pivotal-cf-experimental/cf-probe/big_apps/helpers"
)

type FakeBinarySearchTester struct {
	actual int
}

func NewFakeBinarySearchTester(actual int) *FakeBinarySearchTester {
	return &FakeBinarySearchTester{
		actual: actual,
	}
}

func (tester *FakeBinarySearchTester) Test(value int) bool {
	return value <= tester.actual
}

var _ = Describe("Big Apps Helpers", func() {

	Describe("BinarySearch", func() {
		It("returns a value within the tolerance of the actual value", func() {
			high := 2000
			low := 10
			actual := 500
			tolerance := 5

			tester := NewFakeBinarySearchTester(actual)
			Expect(BinarySearch(tester, low, high, tolerance)).To(BeNumerically("~", actual, tolerance))
		})

		Context("when the actual value is lower than the low limit", func() {
			It("returns the low limit within the tolerance", func() {
				high := 2000
				low := 10
				actual := 2
				tolerance := 5

				tester := NewFakeBinarySearchTester(actual)
				Expect(BinarySearch(tester, low, high, tolerance)).To(BeNumerically("~", low, tolerance))

			})
		})

		Context("when the actual value is greater than the high limit", func() {
			It("returns the high limit within the tolerance", func() {
				high := 2000
				low := 10
				actual := 3000
				tolerance := 5

				tester := NewFakeBinarySearchTester(actual)
				Expect(BinarySearch(tester, low, high, tolerance)).To(BeNumerically("~", high, tolerance))

			})
		})

	})

	Describe("AppSizeBinarySearchTester", func() {

		appPath := "app-path"

		It("Returns true when it can successfully push an app", func() {
			runner := fake_command_runner.New()

			tester := NewAppSizeBinarySearchTester(runner, appPath)

			Expect(tester.Test(100)).To(BeTrue())
		})

		It("Returns false when it can't push an app", func() {
			runner := fake_command_runner.New()

			pushCommand := fake_command_runner.CommandSpec{Path: "gcf"}

			runner.WhenRunning(pushCommand, func(cmd *exec.Cmd) error {
				return errors.New("PUSH FAILED")
			})

			tester := NewAppSizeBinarySearchTester(runner, appPath)

			Expect(tester.Test(100)).To(BeFalse())
		})
	})
})
