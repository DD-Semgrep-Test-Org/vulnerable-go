package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Run()

	// fmt.Printf("%s\n", html.UnescapeString("&nbsp;"))

	print("oh no")
}
