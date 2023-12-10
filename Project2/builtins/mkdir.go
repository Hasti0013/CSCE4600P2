// builtins/mkdir.go

package builtins

import (
	"fmt"
	"os"
)

// Mkdir handles the "mkdir" built-in command.
func Mkdir(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: mkdir <directory>")
	}
	dir := args[0]
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return err
	}
	return nil
}
