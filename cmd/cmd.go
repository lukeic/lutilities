package cmd

import (
	"bytes"
	"os/exec"
	errorUtil "github.com/lukeic/lutilities/error"
)

func Execute(command string, args ...string) (out string, error error) {
	wrappedErr := errorUtil.Wrapped("could not execute command")
	cmd := exec.Command(command, args...)
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut
	err := cmd.Run()
	if err != nil {
		error = wrappedErr(err)
	} else {
		out = cmdOut.String()
	}
	return
}
