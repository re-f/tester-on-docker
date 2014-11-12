package debugLog

import (
	"fmt"
)

func Println(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}
