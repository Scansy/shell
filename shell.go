package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		input, err := reader.ReadString('\n') // Delimiter
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Line read: %s", input)
		// PARSE
		// EXECUTE
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

	if strings.Contains(pwd, homeDir) {
		prompt = hostname + " || ~" + pwd[:len(homeDir)] + "$ "
		return prompt, nil
	}
	prompt = hostname + " || " + pwd + "$ "
	return prompt, nil
}
