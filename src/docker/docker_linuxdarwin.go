// +build !windows

package docker

import (
	"os/exec"
	"strings"
)

var crossCompileCmds = []string{
	"CGO_ENABLED=0",
	"GOOS=" + getImage().os,
	"GOARCH=" + getImage().arch,
	"go test -c -tags inner " + pkname,
}

func executeOnDocker(str string) (string, error) {
	return execute([]string{str})
}

func newCmd(cmds []string) *exec.Cmd {
	cmd := strings.Join(cmds, ";")
	return exec.Command("sh", "-c", cmd)
}
