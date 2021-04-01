package main

import (
	"fmt"
	"oracle/store/keystorage"
)

func inputKey(keystore *keystorage.Keystorage) (err error) {
	fmt.Println("")
	fmt.Println("Please enter the cli/HTTP key, which was provided to you by Oracle")
	fmt.Print("Key: ")
	var key string
	fmt.Scanf("%s\n", &key)
	err = keystore.CheckToken(key)
	if err == nil {
		fmt.Println("Okay, let's continue...")
		return
	} else {
		fmt.Println("I'm not sure I can decrypt your keystore with this key.")
		err = inputKey(keystore)
		return
	}
	return
}

func auth(keystore *keystorage.Keystorage) (err error) {
	fmt.Println("")
	fmt.Println("Let's verify it's you =)")
	err = inputKey(keystore)
	return
}
