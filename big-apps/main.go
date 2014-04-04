package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
)

var bigAppAssets string = "../assets/big-app-base/"

func main() {
  appPath, err := makeBigApp()
  if err != nil {
    fmt.Printf("Making big app failed: %s", err.Error())
  }

  err = pushApp("big-app", appPath)
  if err != nil {
    fmt.Printf("Push failed: %s", err.Error())
  }

}

func pushApp(name string, path string) error {
  cmd := exec.Command("gcf", "push", name, "-p", path)

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  cmd.Start()

  return cmd.Wait()
}

func makeBigApp() (string, error) {
  tempDir, err := ioutil.TempDir("", "big-app")

  if err != nil {
    return "", err
  }

  // debug
  fmt.Printf("Created temp dir %s\n", tempDir)

  cmd := exec.Command("cp", "-r", bigAppAssets, tempDir)

  cmd.Start()

  err = cmd.Wait()

  if err != nil {
    return "", err
  }

  return tempDir, nil
}
