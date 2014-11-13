// +build !inner

package docker

import (
	"fmt"
	"goconf/conf"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func init() {
	filePath, err := searchConfigFile()
	if nil != err {
		panic(err.Error())
	}
	file, err = conf.ReadConfigFile(filePath)
	if nil != err {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}

func RunTestCase(t *testing.T, tc func(t *testing.T)) {
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
	if nil != err {
		fmt.Println("after compile err")
	}
	output, err := runContainer(testInfos[1], filepath.Base(testInfos[0]))
	fmt.Println(output)
	if nil != err {
		t.Fatalf("run container error: " + err.Error())
	}
}
