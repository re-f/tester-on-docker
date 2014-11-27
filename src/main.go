package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("cmd", "/c", "dir & ping")
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	err := cmd.Run()
	if nil != err {
		fmt.Println(err.Error())
	}

	fmt.Println(output.String())
}
