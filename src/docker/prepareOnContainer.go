// +build inner,prepare

package docker

import (
	"fmt"
	"testing"
)

func RunTestCaseDefault(t *testing.T, tc func()) {
	fmt.Println("in 'container, prepare',not run  RunTestCaseDefault")
}

func RunTestCase(t *testing.T, tc func()) {
	fmt.Println("in 'container, prepare',not run  RunTestCase")
}

func Prepare(t *testing.T, tc func(), image *Image) {
	fmt.Println("in 'container, prepare', Preparing")
	tc()
}
