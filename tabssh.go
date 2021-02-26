package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
	"io/ioutil"
	"log"
)

var i int

func eval(expr string) string {
	// FIXME: these need to be unique across runs
	// (or just fix the truncate bug, whatever it is)
	evalpath := fmt.Sprintf("/Users/osnr/t/tabs/last-focused/evals/eval%d.js", i)
	err := ioutil.WriteFile(evalpath, []byte(expr), 0755)
	i = i + 1
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	dat, err := ioutil.ReadFile(evalpath + ".result")
	fmt.Printf("[%s]", dat)
	return string(dat)
}

func url() string {
	url, err := ioutil.ReadFile("/Users/osnr/t/tabs/last-focused/url.txt")
	_ = err
	return string(url[:len(url)-1])
}

func main() {
	ssh.Handle(func(s ssh.Session) {
		for {
			io.WriteString(s, "\u001b[32m")
			io.WriteString(s, url())
			io.WriteString(s, "\u001b[0m> ")

			line := ""
			for {
				ch := make([]byte, 1, 1)
				s.Read(ch)

				line = line + string(ch)
				s.Write(ch)

				// fmt.Println(url.PathEscape(string(ch)))
				if ch[0] == byte('\r') {
					s.Write([]byte{'\n'})
					break
				}
				// FIXME: this doesn't work
				if ch[0] == byte('\b') {
					line = line[:len(line)-1]
				}
			}
			fmt.Printf("Read")
			io.WriteString(s, string(eval(line)))
		}
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
