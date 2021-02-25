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
		err := ioutil.WriteFile("/Users/osnr/t/tabs/last-focused/evals/eval.js", []byte("2+2"), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		}

		dat, err := ioutil.ReadFile("/Users/osnr/t/tabs/last-focused/evals/eval.js.result")

		io.WriteString(s, string(dat))
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
