package docker

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

type TcFunc func(t *testing.T)

var (
	hostPath        = "G:\\Virtual Box\\actiontech-ha"
	boot2dockerPath = "/Users/actiontech-ha"
	containerPath   = "/opt"
)

func runContainer(funcName, pkname string) (string, error) {
	abs, _ := filepath.Abs("./")
	abs = filepath.ToSlash(abs)
	hostPath = filepath.ToSlash(hostPath)
	targetPath := filepath.ToSlash(filepath.Join(strings.Replace(abs, hostPath, boot2dockerPath, 1), pkname+".test"))

	// @todo container name
	runContainer := fmt.Sprintf("sudo docker run -a stdout -i -t --rm=%v -v %v:%v %v %v -test.run=^%v$", true, boot2dockerPath, boot2dockerPath, "ts:base", targetPath, funcName)
	b, err := execute(runContainer)
	return b, err
}

func compileInnerTestCase(pkname string) error {
	cmds := []string{
		"set CGO_ENABLED=0",
		"set GOOS=linux",
		"set GOARCH=amd64",
		"go test -c -tags inner " + pkname,
	}

	fileName := fmt.Sprintf("command_%v.%v", rand.Int(), getScriptSuffix())

	if err := writeFile(fileName, cmds); nil != err {
		return fmt.Errorf("write tmp file error: %v", err)
	}
	defer os.Remove(fileName)

	abs, _ := filepath.Abs("./")
	cmd := exec.Command(filepath.Join(abs, fileName))
	var errput bytes.Buffer
	cmd.Stderr = &errput
	if err := cmd.Run(); nil != err {
		return fmt.Errorf("compile error: %v", errput.String())
	}
	return nil

}

func writeFile(fileName string, lines []string) error {
	// write exec file
	file := strings.Join(lines, getLineEnd())
	if err := ioutil.WriteFile(fileName, []byte(file), 0644); nil != err {
		return err
	}
	return nil
}
