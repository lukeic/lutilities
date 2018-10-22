package cmd

import (
  "bytes"
  "github.com/lukeic/lutilities/pkg/errors"
  "os/exec"
)

func Execute(command string, args ...string) (out string, error error) {
  wrappedErr := errors.Wrapped("could not execute command")
  cmd := exec.Command(command, args...)

  var cmdOut bytes.Buffer
  cmd.Stdout = &cmdOut

  err := cmd.Run()
  if err != nil {
    error = wrappedErr(err)
  }
  out = cmdOut.String()

  return
}
