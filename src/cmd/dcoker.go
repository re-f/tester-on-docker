package cmd

import (
	"bytes"
	"code.google.com/p/go.crypto/ssh"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

type Docker struct {
	ip     string
	port   string
	user   string
	passwd string
}

type TcFunc func(t *testing.T)

var (
	dockerIns = Docker{
		ip:     "192.168.59.103",
		port:   "22",
		user:   "docker",
		passwd: "tcuser",
	}
	hostPath        = "G:\\Virtual Box\\actiontech-ha"
	boot2dockerPath = "/Users/actiontech-ha"
	containerPath   = "/opt"
)

type password string

func (p password) Password(user string) (string, error) {
	return string(p), nil
}

var boot2docker *ssh.ClientConn

func getClient(ip, port, user, passwd string) (*ssh.ClientConn, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.ClientAuth{
			ssh.ClientAuthPassword(password(passwd)),
		},
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", ip, port), config)
	if err != nil {
		return nil, fmt.Errorf("unable to dial remote side:%v", err)
	}
	return client, nil
}
func compileInnerTestCase(pkname string) error {
	cmds := []string{
		"set CGO_ENABLED=0",
		"set GOOS=linux",
		"set GOARCH=amd64",
		"go test -c -tags inner " + pkname,
	}

	// write exec file
	filename := fmt.Sprintf("command_%v.bat", rand.Int())
	batFile := strings.Join(cmds, "\r\n")
	if err := ioutil.WriteFile(filename, []byte(batFile), 0644); nil != err {
		return nil
	}
	defer os.Remove(filename)

	// compile
	var o *bytes.Buffer
	abs, _ := filepath.Abs("./")
	cmd := exec.Command(filepath.Join(abs, filename))
	err := cmd.Run()
	if nil != err {
		return fmt.Errorf("compile test file error : %v; output:%v;", err.Error(), o)
	}
	return nil
}

func runContainer(funcName, pkname string) (string, error) {
	abs, _ := filepath.Abs("./")
	abs = filepath.ToSlash(abs)
	hostPath = filepath.ToSlash(hostPath)
	targetPath := filepath.ToSlash(filepath.Join(strings.Replace(abs, hostPath, boot2dockerPath, 1), pkname+".test"))

	// @todo container name
	runContainer := fmt.Sprintf("sudo docker run -a stdout -i -t --rm=%v -v %v:%v %v %v -test.run=^%v$", true, boot2dockerPath, boot2dockerPath, "ts:base", targetPath, funcName)
	b, err := exec(runContainer)
	return b, err
}

func exec(cmd string) (string, error) {
	client, err := getClient(dockerIns.ip, dockerIns.port, dockerIns.user, dockerIns.passwd)
	// @todo
	if nil != err {
		return "", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if nil != err {
		return "", err
	}
	defer session.Close()

	b, err := session.Output(cmd)
	return stirng(b), err
}
