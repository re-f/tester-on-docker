// +build !inner

package docker

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func RunTestCase(t *testing.T, tc func(t *testing.T)) {
	err := LoadConfig()
	if nil != err {
		t.Fatalf(err.Error())
	}

	pkName, funcName, err := getFuncInfo()
	if nil != err {
		t.Fatalf(err.Error())
	}

	err = compileInnerTestCase(pkName)
	if nil != err {
		t.Fatalf("complie tc error: " + err.Error())
	}

	output, err := runContainer(funcName, filepath.Base(pkName))
	if nil != err {
		t.Fatalf("run container error: %v, output: %v"+err.Error(), output)
	}
}

func getFuncInfo() (pkname, funcName string, err error) {
	funcPc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "", "", fmt.Errorf("get func name and package name error")
	}
	// funcDesc: pkName.funcName
	funcDesc := runtime.FuncForPC(funcPc).Name()
	poitPos := strings.LastIndex(funcDesc, ".") + 1
	return funcDesc[0 : poitPos-1], funcDesc[poitPos:], nil
}
