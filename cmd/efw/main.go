package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sidechannelinc/enclave-code-challenge/pkg/efw"
)

const defaultErrorMessage string = "expected 'status' or 'sync' subcommands"

func main() {
	e := efw.New(context.Background())

	if len(os.Args) < 2 {
		fmt.Println(defaultErrorMessage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "status":
		e.Status()
	case "sync":
		if err := e.Sync(); err != nil {
			fmt.Printf("Error during sync: %v\n", err)
			os.Exit(1)
		}
	case "help":
		fmt.Println("TODO implement a help command")
		fmt.Println(`Available commands:
  efw sync     Fetch and apply firewall rules from remote source
  efw status   Display the currently applied nftables rules
  efw help     Show this help message`)

	default:
		fmt.Println(defaultErrorMessage)
		os.Exit(1)
	}
}
