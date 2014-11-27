// +build !windows

package docker

import (
	"fmt"
	"os/exec"
)

func executeOnDocker(str string) (string, error) {
	return execute(str)
}

func newCmd(cmd ...string) *exec.Cmd {
	c := []string{"/C"}
	c = append(c, cmd...)
	return exec.Command("sh", c...)
}

func getCrossCompileCmd(pkName, os, arch string, isPrepare bool) []string {
	/*var prepare string
	if isPrepare {
		prepare = " prepare"
	} else {
		prepare = ""
	}*/
	cmds := []string{
		"CGO_ENABLED=0",
		fmt.Sprintf("GOOS=%v", os),
		fmt.Sprintf("GOARCH=%v", arch),
		"go", "test", "-c", "-tags", "inner prepare", pkName,
	}
	return cmds
}
