package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "F:/develop/go/src/tester-on-docker/src"
	fmt.Println(filepath.ToSlash(path))

}
