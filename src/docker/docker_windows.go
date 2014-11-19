// +build windows

package docker

import (
	"bytes"
	"code.google.com/p/go.crypto/ssh"
	"fmt"
	"os/exec"
	"strings"
)

type Ssh struct {
	ip     string
	port   string
	user   string
	passwd string
}

func (p Ssh) Password(user string) (string, error) {
	return p.passwd, nil
}

func executeOnDocker(cmd string) (string, error) {
	client, err := getClient()
	if nil != err {
		return "", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if nil != err {
		return "", err
	}
	defer session.Close()

	var output bytes.Buffer
	session.Stderr = &output
	session.Stdout = &output
	err = session.Run(cmd)
	return output.String(), err
}

func getClient() (*ssh.ClientConn, error) {
	config := &ssh.ClientConfig{
		User: getSsh().user,
		Auth: []ssh.ClientAuth{
			ssh.ClientAuthPassword(getSsh()),
		},
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", getSsh().ip, getSsh().port), config)
	if nil != err {
		return nil, fmt.Errorf("unable to dial remote side:%v", err)
	}
	return client, nil
}

func newCmd(cmd string) *exec.Cmd {
	return exec.Command("cmd", "/C", cmd)
}

func getCrossCompileCmd(pkName, os, arch string) string {
	cmds := []string{
		"set CGO_ENABLED=0",
		"set GOOS=" + os,
		"set GOARCH=" + arch,
		"go test -c -tags inner " + pkName,
	}
	return strings.Join(cmds, "&")
}
func getSsh() *Ssh {
	ssh := &Ssh{}
	ssh.ip = getString("ssh", "ip")
	ssh.passwd = getString("ssh", "passwd")
	ssh.port = getString("ssh", "port")
	ssh.user = getString("ssh", "user")
	return ssh
}
