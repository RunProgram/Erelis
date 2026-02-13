package main

import (
	"fmt"
	"os/exec"
)

func main() {
	prompt := "say a funny sentence!"
	out, err := exec.Command("codex", "exec", prompt).Output()
	if err != nil {
		fmt.Println("Error initializing.")
		return
	}
	fmt.Println(out)
}
