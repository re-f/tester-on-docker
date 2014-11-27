package docker

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/ungerik/go-dry/dry"
)

type Image struct {
	Name   string
	Os     string
	Arch   string
	Msg    string
	Author string
}

var compiledPackages = make([]string, 0)

func runContainer(funcName, pkName string, verbose bool, image Image, isRemove, isPrepare bool) (string, string, error) {
	/*
		host path : 	G:/host/path/....
		execution path : 		G:/host/path/src/pkpath
		boot2docker path : 	/docker/path
		=>
		test file path :	/docker/path/src/pkpath
	*/
	if !strings.HasPrefix(getAbs(), getHostPath()) {
		return "", "", fmt.Errorf("must under host path to run test")
	}

	workDir := strings.Replace(getAbs(), getHostPath(), getBoot2DockerPath(), 1)
	testFileName := pkName + ".test"
	testFilePath := filepath.ToSlash(filepath.Join(workDir, testFileName))

	containerName := fmt.Sprintf("%v.%v_%v", pkName, funcName, time.Now().UnixNano())
	var runContainerCmd string
	if isPrepare {
		isRemove = false
		cidfilePath := filepath.ToSlash(filepath.Join(workDir, containerName))
		runContainerCmd = fmt.Sprintf("sudo docker run --name=%v --cidfile=%v -a stdout -i -t --rm=%v -v %v:%v:o -w %v %v %v -test.v=%v -test.run=^%v$", containerName, cidfilePath, isRemove, getBoot2DockerPath(), getBoot2DockerPath(), workDir, image.Name, testFilePath, verbose, funcName)
		debugLog(runContainerCmd)
		output, err := executeOnDocker(runContainerCmd)
		if nil != err {
			return "", output, err
		}
		// get cid
		output, err = executeOnDocker(fmt.Sprintf("cat %v", cidfilePath))
		if nil != err {
			return "", output, err
		} else {
			return output, output, err
		}
	} else {
		runContainerCmd = fmt.Sprintf("sudo docker run --name=%v -a stdout -i -t --rm=%v -v %v:%v:o -w %v %v %v -test.v=%v -test.run=^%v$", containerName, isRemove, getBoot2DockerPath(), getBoot2DockerPath(), workDir, image.Name, testFilePath, verbose, funcName)
		debugLog(runContainerCmd)
		output, err := executeOnDocker(runContainerCmd)
		return "", output, err
	}

}
func removeContainer(cid string) error {
	rmCmd := fmt.Sprintf("docker rm %v", cid)
	_, err := executeOnDocker(rmCmd)
	return err
}
func getAbs() string {
	abs, _ := filepath.Abs("./")
	return filepath.ToSlash(abs)
}

func compileInnerTestCase(pkName string, image Image, isPrepare bool) error {
	if dry.StringInSlice(pkName, compiledPackages) {
		return nil
	}
	cmd := getCrossCompileCmd(pkName, image.Os, image.Arch, isPrepare)
	output, err := execute(cmd...)
	debugLog("comile testcase :%v \n output:%v", cmd, output)
	if nil != err {
		return fmt.Errorf("compile testcase error: %v ,output: %v\n cmd :%v", err.Error(), output, cmd)
	}
	compiledPackages = append(compiledPackages, pkName)
	return nil
}

func execute(strs ...string) (string, error) {
	output, err := newCmd(strs...).CombinedOutput()
	return string(output), err
	/*cmd := newCmd(strs...)
	var output bytes.Buffer
	cmd.Stderr = &output
	cmd.Stdout = &output
	if err := cmd.Start(); nil != err {
		return output.String(), err
	}
	if err := cmd.Wait(); nil != err {
		return output.String(), err
	}
	return output.String(), nil*/
}
