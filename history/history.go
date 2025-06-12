package history

import (
	"fmt"
)

// History stores the command history
type History struct {
	commands []string
}

// Constructor for the History struct
func New() *History {
	return &History{
		commands: make([]string, 0), // Initialize with an empty slice
	}
}

func (h *History) Add(command string) {
	// Add the command to the history
	h.commands = append(h.commands, command)
}

func (h *History) List() {
	if len(h.commands) == 0 {
		fmt.Println("No commands in history.")
		return
	}

	fmt.Println("Command History:")
	for i, cmd := range h.commands {
		fmt.Printf("%d: %s\n", i+1, cmd)
	}
}