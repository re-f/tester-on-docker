// +build inner

package docker

import (
	"testing"
)

func RunTestCase(t *testing.T, tc func(t *testing.T)) {
	tc(t)
}
func RunTestCaseWithPrepare(t *testing.T, imageName string, tc func(t *testing.T)) {
	tc(t)
}
func Prepare(t *testing.T, funcName string, forceNew bool) {
}
