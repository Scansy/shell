package commands

import (
	"os"
)

func cd(args []string) error {
	if len(args) == 0 {
		return os.Chdir(os.Getenv("HOME"))
	}
	
	return os.Chdir(args[0])
}