// +build !windows

package docker

import (
	"os/exec"
	"strings"
)

func executeOnDocker(str string) (string, error) {
	return execute([]string{str})
}

func newCmd(cmds []string) *exec.Cmd {
	cmd := strings.Join(cmds, ";")
	return exec.Command("sh", "-c", cmd)
}
