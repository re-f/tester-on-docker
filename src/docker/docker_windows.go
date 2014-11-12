// +build windows

package docker

import (
	"code.google.com/p/go.crypto/ssh"
	"fmt"
)

type Ssh struct {
	ip     string
	port   string
	user   string
	passwd string
}

func init() {
	dockerIns = getSsh()
}

var (
	dockerIns *Ssh
)

func execute(cmd string) (string, error) {
	client, err := getClient()
	if nil != err {
		return "", err
	}
	// @todo may should mv client.Close() out
	defer client.Close()

	session, err := client.NewSession()
	if nil != err {
		return "", err
	}
	defer session.Close()

	b, err := session.Output(cmd)
	return string(b), err
}

func getClient() (*ssh.ClientConn, error) {
	config := &ssh.ClientConfig{
		User: dockerIns.user,
		Auth: []ssh.ClientAuth{
			ssh.ClientAuthPassword(password(dockerIns.passwd)),
		},
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", dockerIns.ip, dockerIns.port), config)
	if err != nil {
		return nil, fmt.Errorf("unable to dial remote side:%v", err)
	}
	return client, nil
}

type password string

func (p password) Password(user string) (string, error) {
	return string(p), nil
}
func getLineEnd() string {
	return "\r\n"
}
func getScriptSuffix() string {
	return ".bat"
}

func getSsh() *Ssh {
	ssh := &Ssh{}
	ssh.ip = getString("ssh", "ip")
	ssh.passwd = getString("ssh", "passwd")
	ssh.port = getString("ssh", "port")
	ssh.user = getString("ssh", "user")
	return ssh
}
