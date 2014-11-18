package docker

import (
	"bytes"
	"fmt"
	"github.com/ungerik/go-dry/dry"
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
	/*
		host path : 	G:/host/path/....
		run path : 		G:/host/path/src/pkpath
		docker path : 	/docker/path
		-> t
		target path :	/docker/path/src/pkpath

	*/
	abs, _ := filepath.Abs("./")
	targetPath := filepath.ToSlash(filepath.Join(strings.Replace(filepath.ToSlash(abs), getHostPath(), getDockerPath(), 1), pkname+".test"))

	containerName := fmt.Sprintf("%v.%v_%v", pkname, funcName, time.Now().UnixNano())
	runContainerCmd := fmt.Sprintf("sudo docker run --name=%v -a stdout -i -t --rm=%v -v %v:%v:o %v %v -test.run=^%v$", containerName, !isDebug(), getDockerPath(), getDockerPath(), getImage().name, targetPath, funcName)
	debugLog(runContainerCmd)
	b, err := executeOnDocker(runContainerCmd)
	return b, err
}

func compileInnerTestCase(pkname string) error {
	if dry.StringInSlice(pkname, compiledPackages) {
		return nil
	}
	cmds := []string{
		"set CGO_ENABLED=0",
		"set GOOS=" + getImage().os,
		"set GOARCH=" + getImage().arch,
		"go test -c -tags inner " + pkname,
	}

	output, err := execute(cmds)
	if nil != err {
		return fmt.Errorf("compile test case error: %v ,output: %v", err.Error(), output)
	}
	compiledPackages = append(compiledPackages, pkname)
	return nil
}

func execute(strs []string) (string, error) {
	cmd := newCmd(strs)
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
