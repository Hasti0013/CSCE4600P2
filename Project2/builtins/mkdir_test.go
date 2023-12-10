package builtins_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Hasti0013/CSCE4600/Project2/builtins"
)

func TestMkdir(t *testing.T) {
	// Test case for creating a new directory
	t.Run("Create new directory", func(t *testing.T) {
		dir := filepath.Join(t.TempDir(), "newdir")
		err := builtins.Mkdir(dir)
		if err != nil {
			t.Fatalf("Mkdir() error = %v", err)
		}

		_, err = os.Stat(dir)
		if os.IsNotExist(err) {
			t.Errorf("Mkdir() did not create directory")
		}
	})

	// Test case for no arguments
	t.Run("No arguments", func(t *testing.T) {
		err := builtins.Mkdir()
		if err == nil {
			t.Errorf("Mkdir() expected error, got nil")
		}
	})

	// Test case for existing directory
	t.Run("Existing directory", func(t *testing.T) {
		dir := t.TempDir() // This directory already exists
		err := builtins.Mkdir(dir)
		if !os.IsExist(err) {
			t.Errorf("Mkdir() expected os.IsExist error, got %v", err)
		}
	})

	// Test case for invalid directory name
	t.Run("Invalid directory name", func(t *testing.T) {
		dir := filepath.Join(t.TempDir(), "/invalid/name/\000")
		err := builtins.Mkdir(dir)
		if err == nil {
			t.Errorf("Mkdir() expected error, got nil")
		}
	})
}
