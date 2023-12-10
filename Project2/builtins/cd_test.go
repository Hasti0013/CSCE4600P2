package builtins_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/Hasti0013/CSCE4600/Project2/builtins"
)

func TestChangeDirectory(t *testing.T) {
	tmp := t.TempDir()

	type args struct {
		args []string
	}
	tests := []struct {
		name         string
		args         args
		unsetHomedir bool
		wantDir      string
		wantErr      error
		setupFunc    func() string      // optional setup function
		verifyFunc   func(string) error // optional verification function
	}{
		{
			name: "error too many args",
			args: args{
				args: []string{"abc", "def"},
			},
			wantErr: builtins.ErrInvalidArgCount,
		},
		{
			name:    "no args should change to homedir if available",
			wantDir: builtins.HomeDir,
		},
		{
			name:         "no args should error if homedir is blank",
			unsetHomedir: true,
			wantErr:      builtins.ErrInvalidArgCount,
		},
		{
			name: "one arg should change to dir",
			args: args{
				args: []string{tmp},
			},
			wantDir: tmp,
		},
		{
			name: "change to invalid directory",
			args: args{
				args: []string{"/path/to/nonexistent/dir"},
			},
			wantErr: os.ErrNotExist, // or a more specific error if your function returns one
		},
		{
			name: "change to home directory with modified HOME env",
			args: args{
				args: []string{},
			},
			setupFunc: func() string {
				originalHome := os.Getenv("HOME")

				// Create a temporary directory to simulate a fake home directory
				fakeHomeDir := "/tmp/fakehome"
				if err := os.Mkdir(fakeHomeDir, os.ModePerm); err != nil {
					// Handle potential error in creating the directory
					fmt.Printf("Error creating fake home directory: %v\n", err)
				}

				os.Setenv("HOME", fakeHomeDir)
				builtins.HomeDir, _ = os.UserHomeDir()
				return originalHome
			},
			verifyFunc: func(originalHome string) error {
				os.Setenv("HOME", originalHome)
				return nil
			},
			wantDir: "/tmp/fakehome",
		},

		{
			name: "current directory unchanged on error",
			args: args{
				args: []string{"/invalid/dir"},
			},
			setupFunc: func() string {
				origDir, _ := os.Getwd()
				return origDir
			},
			verifyFunc: func(origDir string) error {
				newDir, _ := os.Getwd()
				if origDir != newDir {
					return fmt.Errorf("expected working directory to be unchanged")
				}
				return nil
			},
			wantErr: os.ErrNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup
			var originalState string
			if tt.setupFunc != nil {
				originalState = tt.setupFunc()
			}

			if tt.unsetHomedir {
				oldVal := builtins.HomeDir
				t.Cleanup(func() {
					builtins.HomeDir = oldVal
				})
				builtins.HomeDir = ""
			}

			// testing
			err := builtins.ChangeDirectory(tt.args.args...)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("ChangeDirectory() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else if err != nil {
				t.Fatalf("ChangeDirectory() unexpected error: %v", err)
			} else {
				// verify "happy" path
				wd, err := os.Getwd()
				if err != nil {
					t.Fatalf("Could not get working dir")
				}
				d1, err := os.Stat(wd)
				if err != nil {
					t.Fatalf("Could not stat dir: %v", wd)
				}
				d2, err := os.Stat(tt.wantDir)
				if err != nil {
					t.Fatalf("Could not stat dir: %v", tt.wantDir)
				}
				if !os.SameFile(d1, d2) {
					t.Errorf("Working Directory = %v, wantDir %v", wd, tt.wantDir)
				}
			}

			// cleanup and verify
			if tt.verifyFunc != nil {
				if err := tt.verifyFunc(originalState); err != nil {
					t.Errorf("Verification failed: %v", err)
				}
			}
		})
	}
}
