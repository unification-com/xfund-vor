package main

import (
	"errors"
	"fmt"
	"oracle/store/keystorage"
	"strconv"
)

func noKeyFound(keystorage *keystorage.Keystorage) (fee int64, err error) {
	fmt.Println("")
	fmt.Println("Let's create a new account")
	username, err := promptUntil("Username: ", func(name string) error {
		if name == "" {
			return errors.New("please enter an account username")
		}
		if keystorage.ExistsByUsername(name) {
			return errors.New("this account name is already used")
		}
		return nil
	})
	if err != nil {
		return
	}

	if fee, err = GetFee(); err != nil {
		return
	}

	fmt.Println("")
	fmt.Println("Do you want to add an existing private key or generate a new one?")
	choice, err := promptUntil("[ 1-add existing; 2-generate new ]:	", func(c string) error {
		if c != "1" && c != "2" {
			return errors.New("sorry, I can't understand you: enter 1 or 2")
		}
		return nil
	})
	if err != nil {
		return
	}

	token, err := keystorage.GenerateToken()
	if err != nil {
		return
	}
	if choice == "1" {
		fmt.Println("")
		var private string
		if private, err = prompt("Input your private key (NOTE: it has to start with 0x): "); err != nil {
			return
		}
		if err = keystorage.AddExisting(username, private); err != nil {
			return
		}
		fmt.Println("\nSuccessfully added a private key!")
	} else {
		var keyPrivate string
		if keyPrivate, err = keystorage.GeneratePrivate(username); err != nil {
			return
		}
		fmt.Println("\nSuccessfully generated a private key:")
		fmt.Println(keyPrivate)
	}
	fmt.Print("\nYour daemon api key:   ")
	fmt.Println(token)
	fmt.Println("\nUse this key to login via cli/HTTP (command: oracle-cli settings)")
	fmt.Println("KEEP THIS KEY SAFE! YOU WILL LOSE YOUR KEYSTORAGE DATA WITHOUT IT!")
	fmt.Println("")
	return
}

// GetFee asks for the oracle's fee: a whole number above zero.
func GetFee() (int64, error) {
	fmt.Println("")
	answer, err := promptUntil("Fee: ", func(raw string) error {
		if fee, err := strconv.ParseInt(raw, 10, 64); err != nil || fee <= 0 {
			return errors.New("incorrect amount: enter the fee as a whole number above zero")
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(answer, 10, 64)
}

func FirstRun(keystorage *keystorage.Keystorage) (fee int64, err error) {
	if !stdinIsTerminal() {
		return 0, fmt.Errorf("the keystore has no keys: %w", errNoTerminal)
	}
	fmt.Println("No private key found")
	return noKeyFound(keystorage)
}
