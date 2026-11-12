package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/term"
)

// maxPromptAttempts bounds every interactive question: no answer is ever asked for forever.
const maxPromptAttempts = 5

// input is where prompts read answers from; tests replace it.
var input = bufio.NewReader(os.Stdin)

// errNoTerminal: the oracle needs an answer from the operator but has no terminal to ask on (e.g. it runs under
// systemd), where reading stdin only ever returns end-of-input.
var errNoTerminal = errors.New("this needs an answer from the operator, but the oracle is not running in a terminal")

// stdinIsTerminal is a real TTY check: /dev/null, which systemd usually connects to stdin, is a character
// device too.
func stdinIsTerminal() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}

// prompt prints label and returns the next line of input, trimmed. End of input is an error.
func prompt(label string) (string, error) {
	fmt.Print(label)
	line, err := input.ReadString('\n')
	if err != nil && !(err == io.EOF && line != "") {
		return "", noAnswer(label, err)
	}
	return strings.TrimSpace(line), nil
}

// promptSecret is prompt for a key or passphrase: on a terminal, what is typed isn't echoed.
func promptSecret(label string) (string, error) {
	if !stdinIsTerminal() {
		return prompt(label)
	}
	fmt.Print(label)
	secret, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return "", noAnswer(label, err)
	}
	return strings.TrimSpace(string(secret)), nil
}

func noAnswer(label string, err error) error {
	return fmt.Errorf("no answer to %q: %w", strings.TrimSpace(label), err)
}

// promptUntil asks label until accept returns nil for the answer, at most maxPromptAttempts times, printing why
// each refused answer was refused.
func promptUntil(label string, accept func(answer string) error) (string, error) {
	return askUntil(prompt, label, accept)
}

// promptSecretUntil is promptUntil without echo.
func promptSecretUntil(label string, accept func(answer string) error) (string, error) {
	return askUntil(promptSecret, label, accept)
}

func askUntil(ask func(label string) (string, error), label string, accept func(answer string) error) (string, error) {
	for range maxPromptAttempts {
		answer, err := ask(label)
		if err != nil {
			return "", err
		}
		if err = accept(answer); err == nil {
			return answer, nil
		}
		fmt.Println(err.Error())
	}
	return "", fmt.Errorf("no valid answer to %q after %d attempts", strings.TrimSpace(label), maxPromptAttempts)
}
