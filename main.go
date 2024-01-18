package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "/home")
	cmd.Run()

	fmt.Printf("Now I fix the vuln...")
}
