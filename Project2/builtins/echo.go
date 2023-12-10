// builtins/echo.go

package builtins

import (
	"fmt"
	"io"
	"strings"
)

// Echo handles the "echo" built-in command.
func Echo(w io.Writer, args ...string) error {
	message := strings.Join(args, " ")
	_, err := fmt.Fprintln(w, message)
	return err
}
