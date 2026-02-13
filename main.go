package main

import (
	"fmt"
	"os/exec"
	"os"
	"io"
	"github.com/common-nighthawk/go-figure"
)

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
		fmt.Println("Please install with 'npm install -g @openai/codex'", err)
		return
	}
	fmt.Println(string(out))

	fileContent, err := os.ReadFile("init.txt")
	if err != nil{
		fmt.Println("Unable to reach init file.", err)
		return
	}
	prompt := string(fileContent)
	out, err = exec.Command("codex", "exec", "--skip-git-repo-check", prompt).Output()
	fmt.Println(string(out))
}
