// +build inner

package docker

import (
	"testing"
)

func RunTestCase(t *testing.T, tc func(t *testing.T)) {
	tc(t)
}
