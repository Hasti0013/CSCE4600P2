package builtins_test

import (
	"bytes"
	"testing"

	"github.com/Hasti0013/CSCE4600/Project2/builtins"
)

func TestEcho(t *testing.T) {
	testCases := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Echo with words",
			args: []string{"hello", "world"},
			want: "hello world\n",
		},
		{
			name: "Echo with no arguments",
			args: []string{},
			want: "\n",
		},
		{
			name: "Echo with special characters",
			args: []string{"hello!", "@world$"},
			want: "hello! @world$\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var w bytes.Buffer
			err := builtins.Echo(&w, tc.args...)
			if err != nil {
				t.Fatalf("Echo() error = %v", err)
			}

			if got := w.String(); got != tc.want {
				t.Errorf("Echo() = %v, want %v", got, tc.want)
			}
		})
	}
}
