package windows_demo

import (
	"docker"
	"fmt"

	"testing"
)

func TestDockerTestOnWindows(t *testing.T) {
	docker.RunTestCase(t, func(t *testing.T) {
		fmt.Println("run test...")
	})
}
