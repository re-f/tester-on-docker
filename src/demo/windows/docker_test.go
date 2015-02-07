package windows_demo

import (
	"docker"
	"fmt"
	"runtime"

	"testing"
)

func TestDockerTestOnWindows(t *testing.T) {
	fmt.Println("this test is running on :", runtime.GOOS)
	docker.RunTestCase(t, func(t *testing.T) {
		fmt.Println("this test is executing on :", runtime.GOOS)
		fmt.Println("run test...")
	})
}
