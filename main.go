package main

import (
	_ "embed"
	"fmt"
	"os/exec"

	"github.com/common-nighthawk/go-figure"
)

//go:embed init.txt
var prompt string

func main() {

	start := figure.NewColorFigure("Erelis", "epic", "green", true)
	start.Print()

	println("\033[31m An AI-powered vulnerability hunter written in Go. \033[0m")
	println("\033[31;1m Uses OpenAI's Codex CLI. \033[0m")

	println("\033[31;1m Initializing Codex... \033[0m")

	greeting := "Confirm that Codex is activated by saying a short sentence."

	out, err := exec.Command("codex", "exec", "--skip-git-repo-check", greeting).Output()

	if err != nil {
		fmt.Println("Error starting Codex CLI.", err)
		fmt.Println("Please install with 'npm install -g @openai/codex'")
		return
	}
	fmt.Println(string(out))

	// Use embedded prompt instead of reading file
	out, err = exec.Command("codex", "exec", "--skip-git-repo-check", prompt).Output()
	if err != nil {
		fmt.Println("Scan failed:", err)
		return
	}

	fmt.Println(string(out))
}
