package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	command := os.Getenv("COMMAND")
	pieces := strings.Fields(command)
	cmd := exec.Command(pieces[0], pieces[1:]...)
	cmd.Run()

	fmt.Printf("I just added one line but the vuln is still the same")
}
