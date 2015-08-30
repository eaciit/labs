package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"strings"
	"sync"
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

func TermInOut(w io.Writer, r io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 1)
	out := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for cmd := range in {
			wg.Add(1)
			w.Write([]byte(cmd + "\n"))
			wg.Wait()
		}
	}()
	go func() {
		var (
			buf [1024 * 1024]byte
			t   int
		)
		for {
			n, err := r.Read(buf[t:])
			if err != nil {
				close(in)
				close(out)
				return
			}
			t += n
			if buf[t-2] == '$' {
				out <- string(buf[:t])
				t = 0
				wg.Done()
			}
		}
	}()
	return in, out
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

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if e = s.RequestPty("xterm", 80, 40, modes); e != nil {
		fmt.Println("Unable to start term: " + e.Error())
		return
	}

	commands := []string{"cd /data/goapp/src", "pwd", "ls -al", "cat ~/.bash_profile", "exit"}

	w, _ := s.StdinPipe()
	r, _ := s.StdoutPipe()

	in, out := TermInOut(w, r)
	if e = s.Start("/bin/sh"); e != nil {
		fmt.Println("Unable to start shell: " + e.Error())
		return
	}

	for _, c := range commands {
		in <- c
		outs := strings.Split(<-out, "\n")
		if len(outs) > 1 {
			out := strings.Trim(strings.Join(outs[:len(outs)-1], "\n"), " ")
			fmt.Printf("Output of %s:\n %s\n\n", c, out)
		}
	}
	s.Wait()
}
