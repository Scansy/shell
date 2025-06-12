package commands

import (
	"fmt"
)

func help(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("help: too many arguments, expected 0 or 1")
	} else if len(args) == 0 {
		fmt.Println("Built-in commands at your convenience:")
		fmt.Println("  cd [dir]       - Change the current directory, super useful.")
		fmt.Println("  echo [args...] - Print arguments to standard output, that's probably the terminal.")
		fmt.Println("  exit           - Exit the shell (bye-bye!)")
		fmt.Println("  pwd            - Print current working directory, you can already kinda see it tho.")
		fmt.Println("  clear          - Clear the terminal screen (for when you ls root)")
		fmt.Println("  help           - Show this help message (how did you get here without knowing this?)")
		// Add more as you implement them!
	} else if len(args) == 1 {
		switch args[0] {
		case "cd":
			fmt.Println("cd [dir] - Change the current directory, super useful.")
		case "echo":
			fmt.Println("echo [args...] - Print arguments to standard output, that's probably the terminal.")
		case "exit":
			fmt.Println("exit - Exit the shell (bye-bye!).")
		case "pwd":
			fmt.Println("pwd - Print current working directory, you can already kinda see it tho.")
		case "clear":
			fmt.Println("clear - Clear the terminal screen (for when you ls root)")
		case "help":
			fmt.Println("help - Show this help message (how did you get here without knowing this?)")
		default:
			return fmt.Errorf("help: no help available for '%s'", args[0])
		}
	} 
	return nil
}
