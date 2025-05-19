package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"regexp"
)

// BackgroundProcess is a struct that contains a meta data of a background process.
type BackgroundProcess struct {
	Cmd *exec.Cmd		// Cmd datatype
	Command string		// The command being run
	Input string		// The inputted line to run the command
}

func main() {
	// Variables
	var processes map[string]BackgroundProcess = make(map[string]BackgroundProcess) // input : BackgroundProcess
	var isBackground bool = false
	doneJobs := make(chan string)
	reader := bufio.NewReader(os.Stdin)
	
	// Regex
	backgroundRegex := regexp.MustCompile(`\[(.*?)\]`)


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

		// Joins any goroutines
		select {
			case msg := <-doneJobs:
				fmt.Println("Background job finished: ", msg)
				match := backgroundRegex.FindStringSubmatch(msg)
				if len(match) > 1 { // A sub-match is found
					delete(processes, match[1])
				}

			default: // nothing to join

		}
		fmt.Printf("Num of bg processes: %d\n", len(processes)) // check if bg processes was handled properly
		fmt.Println(processes)
		
		// Skips empty inputs
		if len(input) == 0 {
			continue
		}

		// PARSE
		var args []string = strings.Fields(input) // returns a slice (no size in square brackets)
		fmt.Printf("last arg: %v\n", args[len(args)-1])
		if args[len(args)-1] == "&" {
			isBackground = true
			args = args[:len(args) - 1] // remove ampersand
		}

		// EXECUTE
		var cmd *exec.Cmd = exec.Command(args[0], args[1:]...) // ellipsis syntax to "spread out" the slice
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		if !isBackground {
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			var backgroundProcess BackgroundProcess
			backgroundProcess.Cmd = cmd
			backgroundProcess.Command = args[0]
			backgroundProcess.Input = input
			processes[input] = backgroundProcess
			runBackgroundProcess(backgroundProcess, doneJobs)
			isBackground = false
		}
	}

}

// getPromptInfo creates the prompt line.
func getPromptInfo() (string, error) {
	var prompt string
	hostname, err := os.Hostname()
	if err != nil {
		return prompt, err
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

// runBackgroundProcess launches a goroutine to start and wait for a process in the background
func runBackgroundProcess(bp BackgroundProcess, done chan<- string) {
	go func() {
		err := bp.Cmd.Start()
		if err != nil {
			log.Fatal(err)
		} 
		
		err = bp.Cmd.Wait()
		fmt.Print("Done waiting..")
		if err != nil {
            done <- fmt.Sprintf("Job [%s] exited with error: %v", bp.Input, err)
        } else {
            done <- fmt.Sprintf("Job [%s] completed successfully", bp.Input)
        }

	}()
}