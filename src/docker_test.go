package main

import (
	"docker"
	"fmt"
	"runtime"
	"testing"
)

func TestDocker(t *testing.T) {
	fmt.Println("This test case was execute on ", runtime.GOOS)
	docker.RunTestCase(t, func(t *testing.T) {
		fmt.Println("This test case run on ", runtime.GOOS)
	})
}

func TestDockerWithPrepare(t *testing.T) {
	docker.RunTestCaseWithPrepare(t, "TestDocker", func(t *testing.T) {
		fmt.Println("This test case is base on TestDocker")
	})
}
