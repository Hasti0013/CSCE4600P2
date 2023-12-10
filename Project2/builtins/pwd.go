// builtins/pwd.go

package builtins

import (
	"fmt"
	"io"
	"os"
)

// Pwd handles the "pwd" built-in command.
func Pwd(w io.Writer) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, wd)
	return err
}
