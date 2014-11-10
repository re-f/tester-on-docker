// +build windows

package dockerUnitTester

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CompileInnerTestCase(pkname string) error {
	cmds := []string{
		"set CGO_ENABLED=0",
		"set GOOS=linux",
		"set GOARCH=amd64",
		"go test -c -tags inner " + pkname,
	}
	// write exec file
	filename := fmt.Sprintf("command_%v.bat", rand.Int())
	batFile := strings.Join(cmds, "\r\n")
	if err := ioutil.WriteFile(filename, []byte(batFile), 0644); nil != err {
		return nil
	}
	// exec command
	abs, _ := filepath.Abs("./")
	cmd := exec.Command(filepath.Join(abs, filename))
	err := cmd.Run()
	if nil != err {
		return err
	}
	defer os.Remove(filename)
	return nil
}
