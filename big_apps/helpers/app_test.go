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

			expectedCommand := fake_command_runner.CommandSpec{Path: "gcf", Args: []string{"push", app.Name, "-p", app.Location}}

			Expect(runner).To(HaveExecutedSerially(expectedCommand))
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

			Expect(bigFile.Size()).To(BeNumerically("==", 5*1024*1024))
		})
	})

	Describe("Cleanup", func() {
		It("deletes the local copy of the app", func() {
			runner := command_runner.New(false)
			app, err := NewBigApp(runner, assetPath, 5)
			Expect(err).To(BeNil())

			fileInfo, err := os.Stat(app.Location)
			Expect(err).To(BeNil())
			Expect(fileInfo.IsDir()).To(BeTrue())

			app.Cleanup()

			fileInfo, err = os.Stat(app.Location)
			Expect(err).NotTo(BeNil())
		})

		It("gcf deletes the app", func() {
			runner := fake_command_runner.New()

			app, err := NewBigApp(runner, assetPath, 5)
			Expect(err).To(BeNil())

			app.Cleanup()

			expectedCommand := fake_command_runner.CommandSpec{Path: "gcf", Args: []string{"delete", "-f", app.Name}}

			Expect(runner).To(HaveExecutedSerially(expectedCommand))
		})
	})
})
