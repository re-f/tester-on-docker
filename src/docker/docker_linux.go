// +build !windows

package docker

import (
	"bytes"
	"os/exec"
)

func execute(str string) (string, error) {
	cmd := exec.Command(str)
	var errput bytes.Buffer
	cmd.Stderr = &errput
	if err := cmd.Run(); nil != err {
		return errput.String(), err
	}
	return "", nil
}

func getLineEnd() string {
	return "\n"
}
func getScriptSuffix() string {
	return ".sh"
}
