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

func SendCommand(s *ssh.Session, cmd string) {
	bs, e := s.Output(cmd)
	if e != nil {
		fmt.Println("Fail to send command: " + cmd + " => " + e.Error())
	} else {
		fmt.Println(string(bs))
	}
}

func SendCmd(cmd string, w io.WriteCloser, r io.Reader) {
	var e error
	if _, e = w.Write([]byte(cmd)); e != nil {
		fmt.Printf("Unable to run command %s : %s \n", cmd, e.Error())
		return
	}

	var out []byte
	if _, e = r.Read(out); e != nil {
		fmt.Printf("Unable to read output %s : %s \n", cmd, e.Error())
		return
	} else {
		fmt.Println(string(out))
	}
}

func MuxShell(w io.Writer, r io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 1)
	out := make(chan string, 1)
	var wg sync.WaitGroup
	wg.Add(1) //for the shell itself
	go func() {
		for cmd := range in {
			wg.Add(1)
			w.Write([]byte(cmd + "\n"))
			wg.Wait()
		}
	}()
	go func() {
		var (
			buf [65 * 1024]byte
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
			if buf[t-2] == '$' { //assuming the $PS1 == 'sh-4.3$ '
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
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if e = s.RequestPty("xterm", 80, 40, modes); e != nil {
		fmt.Println("Unable to start term: " + e.Error())
		return
	}

	commands := []string{"cd /data/goapp/src", "pwd", "ls -al", "cat ~/.bash_profile", "exit"}

	w, _ := s.StdinPipe()
	r, _ := s.StdoutPipe()
	/*
		if e = s.Start("/bin/sh"); e != nil {
			fmt.Println("Unable to start shell: " + e.Error())
			return
		}
		for _, c := range commands {
			SendCmd(c, w, r)
		}
		s.Wait()

			s, _ = c.NewSession()
			SendCommand(s, "cd /data/goapp/src")

			s, _ = c.NewSession()
			SendCommand(s, "pwd")
	*/
	//SendCommand(s, "ls -al")

	in, out := MuxShell(w, r)
	if e = s.Start("/bin/sh"); e != nil {
		fmt.Println("Unable to start shell: " + e.Error())
		return
	}

	<-out //ignore the shell output
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
