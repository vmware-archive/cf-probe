package helpers_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry/gunk/command_runner"
	"github.com/cloudfoundry/gunk/command_runner/fake_command_runner"
	. "github.com/cloudfoundry/gunk/command_runner/fake_command_runner/matchers"

	. "github.com/pivotal-cf-experimental/cf-probe/big_apps/helpers"
)

var _ = Describe("Big Apps Helpers", func() {
	assetPath := "../../assets/big-app-base"

	Describe("Push", func() {
		It("calls 'gcf push' with the app name and path", func() {
			runner := fake_command_runner.New()
			app, err := NewBigApp(runner, assetPath, 0)
			Expect(err).To(BeNil())

			app.Push()

			expected_command := fake_command_runner.CommandSpec{Path: "gcf", Args: []string{"push", app.Name, "-p", app.Location}}

			Expect(runner).To(HaveExecutedSerially(expected_command))
		})
	})

	Describe("NewBigApp", func() {
		var app *BigApp

		BeforeEach(func() {
			runner := command_runner.New(false)

			var err error
			app, err = NewBigApp(runner, assetPath, 5)
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			err := os.RemoveAll(app.Location)
			Expect(err).To(BeNil())
		})

		It("copies the app asset", func() {
			files, err := ioutil.ReadDir(app.Location)
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
			bigFile, err := os.Stat(app.Location + "/payload")
			Expect(err).To(BeNil())

			Expect(bigFile.Size()).To(Equal(int64(5 * 1024 * 1024)))
		})
	})
})
