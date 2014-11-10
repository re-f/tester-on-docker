// +build windows

package cmd

import (
	"code.google.com/p/go.crypto/ssh"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Docker struct {
	ip     string
	port   string
	user   string
	passwd string
}

var dockerIns = Docker{
	ip:     "192.168.59.104",
	port:   "22",
	user:   "docker",
	passwd: "tcuser",
}

type password string

func (p password) Password(user string) (string, error) {
	return string(p), nil
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
func Exec(cmd string) {

}
func CompileInnerTestCase(pkname string) error {
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
	// exec command
	abs, _ := filepath.Abs("./")
	cmd := exec.Command(filepath.Join(abs, filename))
	err := cmd.Run()
	if nil != err {
		return err
	}
	defer os.Remove(filename)
	return nil
}

func prepareEnv(t *testing.T) {
	fmt.Println("starting prepare env...")
	cmds := []string{
		"set CGO_ENABLED=0",
		"set GOOS=linux",
		"set GOARCH=amd64",
		"go test -c -tags inner dha/targetService",
	}
	err := execCommands(cmds)
	if nil != err {
		t.Fatalf("complie tc error: " + err.Error())
	}
	// Create a session
	client, err := getClient(dockerIns.ip, dockerIns.port, dockerIns.user, dockerIns.passwd)
	if nil != err {
		t.Fatalf(err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		if nil != err {
			t.Fatalf("unable to create session: %s", err)
		}
	}
	defer session.Close()
	abs, _ := filepath.Abs("./")
	abs = filepath.ToSlash(abs)
	hostPath = filepath.ToSlash(hostPath)
	targetPath := filepath.ToSlash(filepath.Join(strings.Replace(abs, hostPath, boot2dockerPath, 1), "targetService.test"))
	// @todo container name
	runContainer := fmt.Sprintf("sudo docker run -a stdout -i -t --rm=%v -v %v:%v %v %v ", true, targetPath, targetPath, "ts:base", targetPath)
	fmt.Println("run container :", runContainer)
	b, err := session.Output(runContainer)
	fmt.Println("Output: ", string(b))
	if err != nil {
		t.FailNow()
	}
	session.Close()
	client.Close()
}
