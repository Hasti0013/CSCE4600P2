package builtins_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Hasti0013/CSCE4600/Project2/builtins"
)

func TestRemoveFile(t *testing.T) {
	// Create a temporary directory and files for testing
	tempDir := t.TempDir()
	tempFile1 := filepath.Join(tempDir, "tempfile1.txt")
	tempFile2 := filepath.Join(tempDir, "tempfile2.txt")
	for _, file := range []string{tempFile1, tempFile2} {
		if err := os.WriteFile(file, []byte("test"), 0666); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"Remove single file", []string{tempFile1}, false},
		{"Remove multiple files", []string{tempFile1, tempFile2}, false},
		{"Remove non-existent file", []string{"/nonexistentfile"}, true},
		{"Combination of existing and non-existing files", []string{tempFile2, "/nonexistentfile"}, true},
		{"Attempt to remove directory", []string{tempDir}, true},
		{"No arguments", []string{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := builtins.RemoveFile(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
