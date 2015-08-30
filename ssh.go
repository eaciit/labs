package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func Connect() (*ssh.Client, error) {
	cfg := &ssh.ClientConfig{
		User: "developer",
		Auth: []ssh.AuthMethod{
			PublicKeyFile("/users/ariefdarmawan/Dropbox/pvt/aws/developer.pem"),
		},
	}

	client, e := ssh.Dial("tcp", "go.eaciit.com:22", cfg)
	return client, e
}

func SendCommand(s *ssh.Session, cmd string) {
	bs, e := s.Output(cmd)
	if e != nil {
		fmt.Println("Fail to send command: " + cmd + " => " + e.Error())
	} else {
		fmt.Println(string(bs))
	}
}

func main() {
	c, e := Connect()
	if e != nil {
		fmt.Println("Unable to connect: " + e.Error())
	}
	defer c.Close()

	s, e := c.NewSession()
	if e != nil {
		fmt.Println("Unable to start new session: " + e.Error())
		return
	}
	defer s.Close()

	s, _ = c.NewSession()
	SendCommand(s, "cd /data/goapp/src")

	s, _ = c.NewSession()
	SendCommand(s, "pwd")
	//SendCommand(s, "ls -al")
}
