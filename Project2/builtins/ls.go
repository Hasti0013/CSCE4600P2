package builtins

import (
	"fmt"
	"io/ioutil"
)

func ListDirectory(args ...string) error {
	var dir string
	if len(args) == 0 {
		dir = "." // Current directory
	} else if len(args) == 1 {
		dir = args[0]
	} else {
		return fmt.Errorf("ls: too many arguments")
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("ls: %v", err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}
