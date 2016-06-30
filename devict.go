package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(0)
	}

	if os.Args[1] == "events" || os.Args[1] == "e" {
		PrintEvents()
	} else {
		PrintUsage()
	}
}

func PrintUsage() {
	fmt.Println("Usage: devict <command>")
	fmt.Println("               e, events - Lists Events")
	fmt.Print("\n")
}
