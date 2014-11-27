// +build !inner

package docker

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func RunTestCaseDefault(t *testing.T, tc func()) {
	pkName, funcName, err := getFuncInfo(skip + 1)
	if nil != err {
		t.Fatalf(err.Error())
	}
	runTestCase(t, nil, 1, false)
}

func RunTestCase(t *testing.T, tc func(), image *Image) {
	pkName, funcName, err := getFuncInfo(skip + 1)
	if nil != err {
		t.Fatalf(err.Error())
	}
	runTestCase(t, image, 1, false)
}

func Prepare(t *testing.T, tc func(), image *Image) {
	cid := runTestCase(t, tc, image, 1, true)
	err := removeContainer(cid)
	if nil != err {
		t.Fatalf("Prepare error: %v", err.Error())
	}

}

func runTestCase(t *testing.T, pkName, funcName string, image *Image, skip int, isPrepare bool) string {
	err := loadConfig()
	if nil != err {
		t.Fatalf(err.Error())
	}
	if nil == image {
		image = getImage()
	}

	err = compileInnerTestCase(pkName, *image, isPrepare)
	if nil != err {
		t.Fatalf("complie tc error: " + err.Error())
	}

	cid, output, err := runContainer(funcName, filepath.Base(pkName), testing.Verbose(), *image, isDebug(), isPrepare)
	fmt.Println(output)
	if nil != err {
		t.Fatalf("run container error: %v", err.Error())
	}
	return cid
}

func getFuncInfo(skip int) (pkname, funcName string, err error) {
	funcPc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return "", "", fmt.Errorf("get func name and package name error")
	}
	// funcDesc: pkName.funcName
	funcDesc := runtime.FuncForPC(funcPc).Name()
	poitPos := strings.LastIndex(funcDesc, ".") + 1
	return funcDesc[0 : poitPos-1], funcDesc[poitPos:], nil
}
