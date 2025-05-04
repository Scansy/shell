package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		prompt, err := getPromptInfo()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(prompt)

		// READ
		input, err := reader.ReadString('\n') // delimiter
		if err != nil {
			log.Fatal(err)
		}
		
		if len(input) == 0 {
			continue
		}

		// PARSE
		var args []string = strings.Fields(input)// returns a slice (no size in square brackets)

		// EXECUTE
		var cmd *exec.Cmd = exec.Command(args[0], args[1:]...) // ellipsis syntax to "spread out" the slice
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Run()
	}

}

// getPromptInfo creates the prompt line.
func getPromptInfo() (string, error) {
	var prompt string
	hostname, err := os.Hostname()
	if err != nil {
	}

	pwd, err := os.Getwd()
	if err != nil {
		return prompt, err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return prompt, err
	}

	// replace home dir with '~' when appropriate
	if strings.Contains(pwd, homeDir) {
		prompt = hostname + " || ~" + pwd[:len(homeDir)] + "$ "
		return prompt, nil
	}
	prompt = hostname + " || " + pwd + "$ "
	return prompt, nil
}
