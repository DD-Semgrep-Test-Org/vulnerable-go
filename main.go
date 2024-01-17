package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {
	command := os.Getenv("COMMAND")
	pieces := strings.Fields(command)
	cmd := exec.Command(pieces[0], pieces[1:]...)
	cmd.Run()
}
