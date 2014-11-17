package docker

import (
	"fmt"
)

func debug(f func()) {
	if isDebug() {
		f()
	}
}

func notDebug(f func()) {
	if !isDebug() {
		f()
	}
}
func debugLog(msg string, args ...interface{}) {
	if isDebug() {
		fmt.Printf(msg+"\n", args...)
	}
}
