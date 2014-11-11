// +build inner

package cmd

import (
	"testing"
)

func Run(t *testing.T, tc TcFunc) {
	tc(t)
}
