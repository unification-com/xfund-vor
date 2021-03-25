package main

import "fmt"

func noKeyFound() {
	fmt.Print("Do you want to add an existing private key or generate a new one?")
	fmt.Println("[ 1-add existing; 2-generate new ]")
	var addgenerate string
	fmt.Scanln(&addgenerate)
	switch addgenerate {
	case "1":
		//	TODO: add key
		fmt.Print("Successfully added a private key!")
	case "2":
		//	TODO: generate
		fmt.Print("Successfully generated a private key:")
		//  TODO: output private key
		fmt.Print("Your daemon api key:")
		//  TODO: api key
		fmt.Print("Use this key to login via cli/HTTP (command: oracle-cli auth [api key])")
		fmt.Print("KEEP THIS KEYS SAFE!")
	default:
		fmt.Print("Sorry, I can't understand you =(")
		noKeyFound()
	}
}

func firstRun() {
	//	notify that no orivate key found
	fmt.Print("No private key found")
	noKeyFound()

}
