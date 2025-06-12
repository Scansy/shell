package commands
import (
	"fmt"
	"os"
)

func pwd(args []string) error {
	if len(args) >= 1 {
		return fmt.Errorf("pwd: too many arguments")
	}
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("pwd: could not get current directory: %v", err)
	}
	fmt.Println(pwd)
	return nil
}