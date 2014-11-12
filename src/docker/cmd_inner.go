// +build inner

package docker

import (
	"testing"
)

func RunTestCase(t *testing.T, tc TcFunc) {
	tc(t)
}
