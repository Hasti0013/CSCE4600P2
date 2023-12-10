package builtins_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/Hasti0013/CSCE4600/Project2/builtins"
)

func TestListDirectory(t *testing.T) {
	// Create a temporary directory and files for testing
	tempDir := t.TempDir()
	fileNames := []string{"file1.txt", "file2.txt"}
	for _, fname := range fileNames {
		tmpfn := filepath.Join(tempDir, fname)
		if err := ioutil.WriteFile(tmpfn, []byte("test"), 0666); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"List current directory", []string{"."}, false},
		{"List specific directory", []string{tempDir}, false},
		{"Invalid directory", []string{"/invalid/dir"}, true},
		{"Too many arguments", []string{tempDir, tempDir}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := builtins.ListDirectory(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
