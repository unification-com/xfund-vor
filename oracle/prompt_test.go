package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func withInput(t *testing.T, text string) {
	t.Helper()
	original := input
	input = bufio.NewReader(strings.NewReader(text))
	t.Cleanup(func() { input = original })
}

func TestGetFee(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    int64
		wantErr bool
	}{
		{"valid", "100\n", 100, false},
		{"valid without a trailing newline", "250", 250, false},
		{"retries until valid", "abc\n0\n-5\n42\n", 42, false},
		// under systemd stdin is at end-of-input: this used to recurse until the stack overflowed
		{"end of input", "", 0, true},
		{"too many bad answers", strings.Repeat("x\n", maxPromptAttempts+1), 0, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			withInput(t, c.input)
			fee, err := GetFee()
			if c.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, c.want, fee)
		})
	}
}

func TestFirstRunRefusesWithoutTerminal(t *testing.T) {
	if stdinIsTerminal() {
		t.Skip("stdin is a terminal")
	}
	_, err := FirstRun(nil)
	assert.ErrorIs(t, err, errNoTerminal)
}
