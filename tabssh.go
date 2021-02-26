package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
	"io/ioutil"
	"log"
)

func eval(expr string) string {
	err := ioutil.WriteFile("/Users/osnr/t/tabs/last-focused/evals/eval.js", []byte(expr), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	dat, err := ioutil.ReadFile("/Users/osnr/t/tabs/last-focused/evals/eval.js.result")
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
			io.WriteString(s, url())
			io.WriteString(s, "> ")

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
			}
			fmt.Printf("Read")
			io.WriteString(s, string(eval(line)))
		}
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
