package builtins

import (
	"fmt"
	"os"
)

func RemoveFile(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("rm: missing file operand")
	}

	for _, file := range args {
		err := os.Remove(file)
		if err != nil {
			return fmt.Errorf("rm: %v", err)
		}
	}
	return nil
}
