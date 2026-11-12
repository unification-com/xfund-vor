package main

import (
	"errors"
	"fmt"
	"oracle/store/keystorage"
	"strconv"
)

// noKeyFound creates the oracle's first key under a new random keystore key, which it prints (the only copy),
// and an HTTP API key, which it writes where the admin commands read it.
func noKeyFound(keystore *keystorage.Keystorage) (fee int64, token string, err error) {
	fmt.Println("")
	fmt.Println("Let's create a new account")
	username, err := promptUntil("Username: ", func(name string) error {
		if name == "" {
			return errors.New("please enter an account username")
		}
		if keystore.ExistsByUsername(name) {
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

	if token, err = keystore.GenerateToken(); err != nil {
		return
	}
	if choice == "1" {
		fmt.Println("")
		var private string
		private, err = promptSecretUntil("Input your private key (hex; not shown as you type): ", func(key string) error {
			_, err := keystorage.ParsePrivateKey(key)
			return err
		})
		if err != nil {
			return
		}
		if err = keystore.AddExisting(username, private); err != nil {
			return
		}
		fmt.Println("\nSuccessfully added a private key!")
	} else {
		var address string
		if address, err = keystore.GeneratePrivate(username); err != nil {
			return
		}
		fmt.Println("\nSuccessfully generated a private key for address", address)
		fmt.Println("It exists only in the keystore: back up the keystore file and keep the key below safe.")
	}
	if err = issueAPIKey(keystore, defaultAPIKeyFile()); err != nil {
		return
	}
	fmt.Print("\nYour keystore key:   ")
	fmt.Println(token)
	fmt.Println("Start the oracle with it (-k). KEEP THIS KEY SAFE! YOU WILL LOSE YOUR KEYSTORAGE DATA WITHOUT IT!")
	fmt.Println("\nThe admin commands' API key is in", defaultAPIKeyFile())
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

// FirstRun creates the oracle's first key and returns its fee and the key that unlocks the keystore.
func FirstRun(keystore *keystorage.Keystorage) (fee int64, token string, err error) {
	if !stdinIsTerminal() {
		return 0, "", fmt.Errorf("the keystore has no keys: %w", errNoTerminal)
	}
	fmt.Println("No private key found")
	return noKeyFound(keystore)
}
