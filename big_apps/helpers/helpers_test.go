package helpers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/pivotal-cf-experimental/cf-probe/big_apps/helpers"

	"io/ioutil"
	"os"

	"github.com/cloudfoundry/gunk/command_runner"
	"github.com/cloudfoundry/gunk/command_runner/fake_command_runner"
	. "github.com/cloudfoundry/gunk/command_runner/fake_command_runner/matchers"
)

var _ = Describe("Big Apps Helpers", func() {
	Describe("PushApp", func() {
		It("calls 'gcf push' with the app name and path", func() {
			runner := fake_command_runner.New()

			PushApp(runner, "app-name", "app-path")

			expected_command := fake_command_runner.CommandSpec{Path: "gcf", Args: []string{"push", "app-name", "-p", "app-path"}}

			Expect(runner).To(HaveExecutedSerially(expected_command))
		})
	})

	Describe("MakeBigApp", func() {
		assetPath := "../../assets/big-app-base"
		var location string

		BeforeEach(func() {
			runner := command_runner.New(false)

			var err error
			location, err = MakeBigApp(runner, assetPath, 5)
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			err := os.RemoveAll(location)
			Expect(err).To(BeNil())
		})

		It("copies the app asset", func() {
			files, err := ioutil.ReadDir(location)
			Expect(err).To(BeNil())

			fileNames := []string{}
			for _, fileInfo := range files {
				fileNames = append(fileNames, fileInfo.Name())
			}

			Expect(fileNames).To(ContainElement("Gemfile"))
			Expect(fileNames).To(ContainElement("vendor"))
			Expect(fileNames).To(ContainElement("Gemfile.lock"))
			Expect(fileNames).To(ContainElement("config.ru"))
		})

		It("adds a file of the given number of megabytes", func() {
			bigFile, err := os.Stat(location + "/payload")
			Expect(err).To(BeNil())

			Expect(bigFile.Size()).To(Equal(int64(5 * 1024 * 1024)))
		})
	})
})
