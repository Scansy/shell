package commands

import (
	"os"
	"fmt"
)

func cd(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cd: missing argument")
	}
	
	return os.Chdir(args[0])
}