package main

import (
	"fmt"
	"oracle/store/keystorage"
	"strconv"
)

func noKeyFound(keystorage *keystorage.Keystorage) (fee int64, paysGas bool, err error) {
	var token string

	fmt.Println("")
	fmt.Println("Let's create a new account")
	fmt.Print("Username: ")
	var addusername string
	fmt.Scanf("%s\n", &addusername)

	fee, err = GetFee()
	paysGas, err = GetProviderPaysGas()

	if keystorage.ExistsByUsername(addusername) {
		fmt.Println("This account name is already used")
		noKeyFound(keystorage)
	} else if addusername == "" {
		fmt.Println("Please enter account username.")
		noKeyFound(keystorage)
	}
	fmt.Println("")
	fmt.Println("Do you want to add an existing private key or generate a new one?")
	fmt.Print("[ 1-add existing; 2-generate new ]:	")
	var addgenerate string
	fmt.Scanf("%s\n", &addgenerate)

	switch addgenerate {
	case "1":
		token, err = keystorage.GenerateToken()
		fmt.Println("")
		fmt.Print("Input your private key (NOTE: it has to start with 0x): ")
		var addprivate string
		fmt.Scanf("%s\n", &addprivate)
		err = keystorage.AddExisting(addusername, addprivate)
		if err != nil {
			return
		}
		fmt.Println("\nSuccessfully added a private key!\n")
		fmt.Print("Your daemon api key:   ")
		fmt.Println(token)
		fmt.Println("\nUse this key to login via cli/HTTP (command: oracle-cli settings)")
		fmt.Println("KEEP THIS KEY SAFE! YOU WILL LOSE YOUR KEYSTORAGE DATA WITHOUT IT!")
		fmt.Println("")
	case "2":
		token, err = keystorage.GenerateToken()
		if err != nil {
			return
		}
		var keyPrivate string
		keyPrivate, err = keystorage.GeneratePrivate(addusername)
		if err != nil {
			return fee, paysGas, err
		}
		fmt.Println("\nSuccessfully generated a private key:")
		fmt.Println(keyPrivate)
		fmt.Print("\nYour daemon api key:   ")
		fmt.Println(token)
		fmt.Println("\nUse this key to login via cli/HTTP (command: oracle-cli settings)")
		fmt.Println("KEEP THIS KEYS SAFE!")
	default:
		fmt.Println("Sorry, I can't understand you =(")
		fee, paysGas, err = noKeyFound(keystorage)
	}
	return
}

func GetFee() (input int64, err error) {
	var rawInput string
	fmt.Println("")
	fmt.Print("Fee: ")
	_, err = fmt.Scanf("%s\n", &rawInput)
	input, err = strconv.ParseInt(rawInput, 10, 64)
	if err != nil {
		fmt.Println("Incorrect amount")
		input, err = GetFee()
	}
	if input == 0 {
		fmt.Println("Please enter Fee.")
		input, err = GetFee()
	}
	return
}

func GetProviderPaysGas() (input bool, err error) {
	var rawInput string
	fmt.Println("")
	fmt.Print("Does provider pay gas? [true|false]: ")
	_, err = fmt.Scanf("%s\n", &rawInput)
	input, err = strconv.ParseBool(rawInput)
	if err != nil {
		fmt.Println("Incorrect provider pays gas parameter")
		input, err = GetProviderPaysGas()
	}

	return
}

func FirstRun(keystorage *keystorage.Keystorage) (fee int64, paysGas bool, err error) {
	//	notify that no orivate key found
	fmt.Println("No private key found")
	fee, paysGas, err = noKeyFound(keystorage)
	return
}
