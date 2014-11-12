// +build !inner

package docker

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func RunTestCase(t *testing.T, tc TcFunc) {
	funcPc, _, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatalf("get func name and package name error")
	}
	testInfos := strings.Split(runtime.FuncForPC(funcPc).Name(), ".")
	err := compileInnerTestCase(testInfos[0])
	if nil != err {
		t.Fatalf("complie tc error: " + err.Error())
		return
	}
	output, err := runContainer(testInfos[1], filepath.Base(testInfos[0]))
	fmt.Println(output)
	if nil != err {
		t.Fatalf("run container error: " + err.Error())
	}
}
