package builtins_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/Hasti0013/CSCE4600/Project2/builtins"
)

func TestPwd(t *testing.T) {
	// Check for CI environment variable
	if _, ci := os.LookupEnv("CI"); ci {
		t.Skip("Skipping TestPwd in CI environment")
	}

	var w bytes.Buffer
	err := builtins.Pwd(&w)
	if err != nil {
		t.Fatalf("Pwd() error = %v", err)
	}

	// Validate the output as an existing directory
	got := strings.TrimSpace(w.String())
	if _, err := os.Stat(got); os.IsNotExist(err) {
		t.Errorf("Pwd() returned a non-existent directory path: %v", got)
	}
}

func TestPwdConsistency(t *testing.T) {
	// Check for CI environment variable
	if _, ci := os.LookupEnv("CI"); ci {
		t.Skip("Skipping TestPwd in CI environment")
	}

	var w1, w2 bytes.Buffer

	// Call Pwd twice
	err1 := builtins.Pwd(&w1)
	err2 := builtins.Pwd(&w2)

	if err1 != nil {
		t.Fatalf("Pwd() (1st call) error = %v", err1)
	}

	if err2 != nil {
		t.Fatalf("Pwd() (2nd call) error = %v", err2)
	}

	// Check if the output is consistent
	got1 := strings.TrimSpace(w1.String())
	got2 := strings.TrimSpace(w2.String())

	if got1 != got2 {
		t.Errorf("Pwd() returned inconsistent directory paths:\n1st call: %v\n2nd call: %v", got1, got2)
	}
}
