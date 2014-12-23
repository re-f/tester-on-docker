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
var repository = "auto_by_tester"

func runContainer(funcName, pkName string, im image, verbose, isRemove, isPrepare bool) (string, string, error) {
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
		runContainerCmd = fmt.Sprintf("docker run --name=%v --cidfile=%v -a stdout -i -t --rm=%v -v %v:%v:o -w %v %v %v -test.v=%v -test.run=^%v$ ", containerName, cidfilePath, isRemove, getBoot2DockerPath(), getBoot2DockerPath(), workDir, im.name, testFilePath, verbose, funcName)
		debugLog("[run docker container]%v", runContainerCmd)
		prepareOutput, err := executeOnDocker(runContainerCmd)
		if nil != err {
			return "", prepareOutput, err
		}
		// get cid
		output, err := executeOnDocker(fmt.Sprintf("cat %v", cidfilePath))
		if nil != err {
			return output, prepareOutput, err
		} else {
			return output, prepareOutput, err
		}
	} else {
		runContainerCmd = fmt.Sprintf("sudo docker run --name=%v -a stdout -i -t --rm=%v -v %v:%v:o -w %v %v %v -test.v=%v -test.run=^%v$", containerName, isRemove, getBoot2DockerPath(), getBoot2DockerPath(), workDir, im.name, testFilePath, verbose, funcName)
		debugLog("[Info]%v", runContainerCmd)
		output, err := executeOnDocker(runContainerCmd)
		return "", output, err
	}
}

func getAbs() string {
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

func removeContainer(cid string) error {
	rmCmd := fmt.Sprintf("docker rm %v", cid)
	output, err := executeOnDocker(rmCmd)
	if nil != err {
		return fmt.Errorf("remove container error: %v, output: %v", err.Error(), output)
	}
	return nil
}

func removeImage(name string) error {
	rmCmd := fmt.Sprintf("docker rmi %v", name)
	output, err := executeOnDocker(rmCmd)
	if nil != err {
		return fmt.Errorf("remove container error: %v, output: %v", err.Error(), output)
	}
	return nil
}

func buildImage(cid string, imageName string) error {
	output, err := executeOnDocker(fmt.Sprintf("docker  commit -a \"build by tester_on_docker\" -m \"auto\"  %v %v:%v", cid, repository, imageName))
	if nil != err {
		return fmt.Errorf("error: %v ,info: %v", err.Error(), output)
	}
	return nil

}
func isImageExist(repository, tag string) bool {
	output, err := executeOnDocker(fmt.Sprintf("docker images | grep %v", repository))
	if nil != err {
		return false
	}
	return strings.Contains(output, tag)
}
