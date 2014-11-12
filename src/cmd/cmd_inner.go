// +build inner

package cmd

import (
	"testing"
)

func RunTestCase(t *testing.T, tc TcFunc) {
	tc(t)
}
