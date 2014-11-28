// +build container,prepare

package docker

import (
	"testing"
)

func RunTestCaseDefault(t *testing.T, tc func()) {
}

func RunTestCase(t *testing.T, tc func()) {
}

func Prepare(t *testing.T, tc func(), image *Image, forceNew bool) {
	tc()
}
