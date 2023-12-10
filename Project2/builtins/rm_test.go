package builtins_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Hasti0013/CSCE4600/Project2/builtins"
)

func TestRemoveFile(t *testing.T) {
	// Create a temporary file for testing
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "tempfile.txt")
	if err := os.WriteFile(tempFile, []byte("test"), 0666); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"Remove single file", []string{tempFile}, false},
		{"Remove non-existent file", []string{"/nonexistentfile"}, true},
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
