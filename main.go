package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
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
	doneJobs := make(chan error)
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
		input = strings.TrimSpace(input) // remove trailing newline

		// Joins any goroutines
		for {
			select {
				case errMsg := <-doneJobs:
					match := backgroundRegex.FindStringSubmatch(errMsg.Error())
					if len(match) > 1 { // A sub-match is found
						delete(processes, match[1])
					} 

				default: // nothing to join
					goto continueREPL // go to continueREPL label
			}
		}
		continueREPL:
		fmt.Println("Processes: ", processes)
		
		// Skips empty inputs
		if len(input) == 0 {
			continue
		}

		// PARSE
		var args []string = strings.Fields(input) // returns a slice (no size in square brackets)
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
				fmt.Println("shell: unable to run command")
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
func runBackgroundProcess(bp BackgroundProcess, done chan<- error) {
	go func()  {
		err := bp.Cmd.Start()
		if err != nil {
			done <- errors.New("shell: unable to start command")
		} 
		
		err = bp.Cmd.Wait()
		fmt.Print("Done waiting..")
		if err != nil {
			done <- errors.New("shell: unable to wait for command")
		} else { // return the inputted line
			done <- errors.New("[" + bp.Input + "]")
		}
	}()
}