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
	"time"
)

type image struct {
	name string
	os   string
	arch string
}

var compiledPackages = make([]string, 0)

func runContainer(funcName, pkname string) (string, error) {
	abs, _ := filepath.Abs("./")
	abs = filepath.ToSlash(abs)
	hostPath := filepath.ToSlash(getHostPath())
	targetPath := filepath.ToSlash(filepath.Join(strings.Replace(abs, hostPath, getDockerPath(), 1), pkname+".test"))
	containerName := fmt.Sprintf("%v.%v_%v", pkname, funcName, time.Now().Unix())
	runContainerCmd := fmt.Sprintf("sudo docker run --name=%v -a stdout -i -t --rm=%v -v %v:%v:o %v %v -test.run=^%v$", containerName, !isDebug(), getDockerPath(), getDockerPath(), getImage().name, targetPath, funcName)
	debugLog(runContainerCmd)
	b, err := executeOnDocker(runContainerCmd)
	return b, err
}

func compileInnerTestCase(pkname string) error {
	for _, compiledPackage := range compiledPackages {
		if pkname == compiledPackage {
			return nil
		}
	}
	cmds := []string{
		"set CGO_ENABLED=0",
		"set GOOS=" + getImage().os,
		"set GOARCH=" + getImage().arch,
		"go test -c -tags inner " + pkname,
	}

	fileName := fmt.Sprintf("command_%v.%v", rand.Int(), getScriptSuffix())

	if err := writeFile(fileName, cmds); nil != err {
		return fmt.Errorf("write tmp file error: %v", err.Error())
	}
	defer notDebug(func() { os.Remove(fileName) })

	abs, _ := filepath.Abs("./")
	output, err := execute(filepath.Join(abs, fileName))
	if nil != err {
		return fmt.Errorf("compile test case error: %v ,output: %v", err.Error(), output)
	}
	compiledPackages = append(compiledPackages, pkname)
	return nil
}

func writeFile(fileName string, lines []string) error {
	file := strings.Join(lines, getLineEnd())
	if err := ioutil.WriteFile(fileName, []byte(file), 0644); nil != err {
		return err
	}
	return nil
}

func execute(str string) (string, error) {
	cmd := exec.Command(str)
	var output bytes.Buffer
	cmd.Stderr = &output
	cmd.Stdout = &output
	if err := cmd.Start(); nil != err {
		return output.String(), err
	}
	if err := cmd.Wait(); nil != err {
		return output.String(), err
	}

	return output.String(), nil
}
