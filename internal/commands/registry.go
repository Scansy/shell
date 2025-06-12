package commands

// Registry is simply a map of command names to their respective functions.
var Registry = map[string]func(args []string) error{
	"cd":   cd,
	"echo": echo,
	"exit": exit,
	"pwd":  pwd,
	"help": help,
	// Add more commands here as needed
}