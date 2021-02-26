package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		for {
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

			err := ioutil.WriteFile("/Users/osnr/t/tabs/last-focused/evals/eval.js", []byte(line), 0755)
			if err != nil {
				fmt.Printf("Unable to write file: %v", err)
			}

			dat, err := ioutil.ReadFile("/Users/osnr/t/tabs/last-focused/evals/eval.js.result")
			fmt.Printf("[%s]", dat)
			io.WriteString(s, string(dat))
		}
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
