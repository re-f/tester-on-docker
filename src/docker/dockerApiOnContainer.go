// +build container,!prepare

package docker

import (
	"testing"
)

func RunTestCaseDefault(t *testing.T, tc func()) {
	tc()
}

func RunTestCase(t *testing.T, tc func()) {
	tc()
}

func Prepare(t *testing.T, tc func(), image *Image, forceNew bool) {
}
