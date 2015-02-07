// +build inner

package docker

import (
	"testing"
)

func RunTestCase(t *testing.T, tc func(t *testing.T)) {
	tc(t)
}
func RunTestCaseWithPrepare(t *testing.T, funcName string, tc func(t *testing.T)) {
	tc(t)
}
func prepare(t *testing.T, funcName string) {
}
