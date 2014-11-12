// +build windows

package docker

import (
	"code.google.com/p/go.crypto/ssh"
	"fmt"
)

type Docker struct {
	ip     string
	port   string
	user   string
	passwd string
}

var (
	boot2docker *ssh.ClientConn
	dockerIns   = Docker{
		ip:     "192.168.59.103",
		port:   "22",
		user:   "docker",
		passwd: "tcuser",
	}
)

func execute(cmd string) (string, error) {
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
	return string(b), err
}

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
