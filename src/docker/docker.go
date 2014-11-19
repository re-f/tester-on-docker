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

func runContainer(funcName, pkName string,verbose bool) (string, error) {
	/*
		host path : 	G:/host/path/....
		execution path : 		G:/host/path/src/pkpath
		boot2docker path : 	/docker/path
		=> 
		test file path :	/docker/path/src/pkpath
	*/
	if !strings.HasPrefix(getAbs(),getHostPath() ) {
		return "",fmt.Errorf("must under host path to run test")
	}
	testFileName :=pkName+".test"
	testFilePath := filepath.ToSlash( filepath.Join(strings.Replace(getAbs(), getHostPath(), getBoot2DockerPath(), 1), testFileName))

	containerName := fmt.Sprintf("%v.%v_%v", pkName, funcName, time.Now().UnixNano())
	runContainerCmd := fmt.Sprintf("sudo docker run --name=%v -a stdout -i -t --rm=true -v %v:%v:o %v %v -test.v=%v -test.run=^%v$", containerName,  getBoot2DockerPath(), getBoot2DockerPath(), getImage().name, testFilePath,verbose, funcName)
	debugLog(runContainerCmd)
	return executeOnDocker(runContainerCmd)
}

func getAbs()string {
	abs, _ := filepath.Abs("./")
	return filepath.ToSlash(abs)
}

func compileInnerTestCase(pkName string) error {
	if dry.StringInSlice(pkName, compiledPackages) {
		return nil
	}
	output, err := execute(getCrossCompileCmd(pkName, getImage().os, getImage().arch))
	if nil != err {
		return fmt.Errorf("compile test case error: %v ,output: %v", err.Error(), output)
	}
	compiledPackages = append(compiledPackages, pkName)
	return nil
}

func execute(strs string) (string, error) {
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
