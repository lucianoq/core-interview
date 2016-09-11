package main

import (
	"flag"
	"os"
	"fmt"
	httpclient "core-interview/client"
)

func main() {
	var idFlag, dataFlag, keyFlag string

	storeCommand := flag.NewFlagSet("store", flag.ExitOnError)
	storeCommand.StringVar(&idFlag, "id", "", "Id of the message to be stored")
	storeCommand.StringVar(&dataFlag, "data", "", "Message to be stored")

	retrieveCommand := flag.NewFlagSet("retrieve", flag.ExitOnError)
	retrieveCommand.StringVar(&idFlag, "id", "", "Id of the message to be retrieved")
	retrieveCommand.StringVar(&keyFlag, "key", "", "AES key to use to decrypt")

	if len(os.Args) == 1 {
		fmt.Println("usage: yoti <command> [<args>]")
		fmt.Println("The most commonly used yoti commands are: ")
		fmt.Println(" store       Store an encrypted message")
		fmt.Println(" retrieve    Retrieve and decrypted message")
		return
	}

	switch os.Args[1] {
	case "store":
		if err := storeCommand.Parse(os.Args[2:]); err != nil {
			fmt.Printf("Error on parsing command, %s\n", err.Error())
			return
		}
		if idFlag == "" {
			fmt.Println("Please supply the id using --id option.")
			return
		}
		if dataFlag == "" {
			fmt.Println("Please supply the data using --data option.")
			return
		}
		var httpC httpclient.HttpClient
		key, err := httpC.Store([]byte(idFlag), []byte(dataFlag))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(string(key))
	case "retrieve":
		if err := retrieveCommand.Parse(os.Args[2:]); err != nil {
			fmt.Printf("Error on parsing command, %s\n", err.Error())
			return
		}
		if idFlag == "" {
			fmt.Println("Please supply the id using --id option.")
			return
		}
		if keyFlag == "" {
			fmt.Println("Please supply the key using --key option.")
			return
		}
		var httpC httpclient.HttpClient
		plaintext, err := httpC.Retrieve([]byte(idFlag), []byte(keyFlag))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(string(plaintext))
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
