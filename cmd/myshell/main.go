package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		fmt.Println("error when trying to read command")
	} else {
		command = strings.TrimSpace(command)
		fmt.Println(command + ": command not found")
	}
}
