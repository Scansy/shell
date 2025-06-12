package commands

import (
	"fmt"
	"os"
)

func exit(args []string) error {
	// Exit the shell
	fmt.Println("Goodbye!") // TODO: Add random, funny messages
	os.Exit(0)
	return nil // Never reached, included to satisfy the function signature
}