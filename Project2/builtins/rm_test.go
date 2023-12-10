package builtins_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Hasti0013/CSCE4600/Project2/builtins"
)

func TestRemoveFile(t *testing.T) {
func RemoveFile(args ...string) error {
    if len(args) < 1 {
        return fmt.Errorf("rm: missing file operand")
    }

    for _, file := range args {
        fileInfo, err := os.Stat(file)
        if err != nil {
            return fmt.Errorf("rm: %v", err)
        }

        if fileInfo.IsDir() {
            return fmt.Errorf("rm: %s is a directory", file)
        }

        err = os.Remove(file)
        if err != nil {
            return fmt.Errorf("rm: %v", err)
        }
    }
    return nil
}
{
		{
			name: "Remove single file",
			setupFunc: func() []string {
				return []string{createTempFile(t, tempDir, "tempfile1.txt")}
			},
			wantErr: false,
		},
		{
			name: "Remove multiple files",
			setupFunc: func() []string {
				return []string{
					createTempFile(t, tempDir, "tempfile2.txt"),
					createTempFile(t, tempDir, "tempfile3.txt"),
				}
			},
			wantErr: false,
		},
		{
			name: "Remove non-existent file",
			args: []string{"/nonexistentfile"},
			wantErr: true,
		},
		{
			name: "Attempt to remove directory",
			args: []string{tempDir},
			wantErr: true,
		},
		{
			name: "No arguments",
			args: []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			if tt.setupFunc != nil {
				tt.args = tt.setupFunc()
			}

			// testing
			err := builtins.RemoveFile(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// createTempFile is a helper function to create a temporary file
func createTempFile(t *testing.T, dir, filename string) string {
	path := filepath.Join(dir, filename)
	if err := os.WriteFile(path, []byte("test"), 0666); err != nil {
		t.Fatalf("failed to create temporary file: %s", err)
	}
	return path
}
