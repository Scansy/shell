package commands

import (
	"fmt"
	"strings"
)

func echo(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("echo: no arguments provided")
	}

	// Join the arguments with a space and print them
	fmt.Println(strings.Join(args, " "))
	return nil
}