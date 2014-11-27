// +build inner,!prepare

package docker

import (
	"fmt"
	"testing"
)

func RunTestCaseDefault(t *testing.T, tc func()) {
	fmt.Println("in 'container, !prepare' RunTestCaseDefault")
	tc()
}

func RunTestCase(t *testing.T, tc func()) {
	fmt.Println("in 'container, !prepare' RunTestCase")
	tc()
}

func Prepare(t *testing.T, tc func(), image *Image) {
	fmt.Println("in 'container, !prepare' not run Prepare")
}
