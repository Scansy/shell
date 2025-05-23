package commands

import (
	"fmt"
	"os"
)

func exit() {
	// Exit the shell
	fmt.Println("Goodbye!") // TODO: Add random, funny messages
	os.Exit(0)
}