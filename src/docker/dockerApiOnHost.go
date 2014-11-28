// +build !container

package docker

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func RunTestCaseDefault(t *testing.T, tc func()) {
	runTestCase(t, nil, 1, false)
}

func RunTestCase(t *testing.T, tc func(), image *Image) {
	runTestCase(t, image, 1, false)
}

func Prepare(t *testing.T, tc func(), image *Image, forceNew bool) {
	cid := runTestCase(t, nil, 1, true)

	if err := buildImage(cid, image); nil != err {
		t.Fatalf("Prepare error: %v", err.Error())
	}

	if err := removeContainer(cid); nil != err {
		t.Fatalf("Prepare error: %v", err.Error())
	}

}

func runTestCase(t *testing.T, image *Image, skip int, isPrepare bool) string {
	err := loadConfig()
	if nil != err {
		t.Fatalf(err.Error())
	}
	if nil == image {
		image = getImage()
	}
	pkName, funcName, err := getFuncInfo(skip + 1)
	if nil != err {
		t.Fatalf(err.Error())
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
