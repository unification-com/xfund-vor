package main

import (
	"fmt"
	"oracle/store/keystorage"
)

func inputKey(keystorage *keystorage.Keystorage) (err error) {
	fmt.Println("")
	fmt.Println("Please enter the cli/HTTP key, which was provided to you by Oracle")
	fmt.Print("Key: ")
	var key string
	fmt.Scanf("%s\n", &key)
	err = keystorage.CheckToken(key)
	if err == nil {
		fmt.Println("Okay, let's continue...")
	} else {
		fmt.Println("I'm not sure I can decrypt your keystore with this key.")
		err = inputKey(keystorage)
	}
	return
}

func auth(keystorage *keystorage.Keystorage) (err error) {
	fmt.Println("")
	fmt.Println("Let's verify it's you =)")
	err = inputKey(keystorage)
	return
}
