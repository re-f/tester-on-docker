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

// @todo
func init() {
	filePath, err := searchConfigFile()
	if nil != err {
		panic(err.Error())
	}
	file, err = conf.ReadConfigFile(filePath)
	debugLog("config at %v", filePath)
	if nil != err {
		fmt.Println(err.Error())
		panic(err.Error())
	}
}

func RunTestCase(t *testing.T, tc func(t *testing.T)) {

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
		t.Fatalf("run container error: ,output: %v"+err.Error(), output)
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
