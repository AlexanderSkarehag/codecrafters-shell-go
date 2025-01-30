package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func checkIfBuiltin(t string) {
	switch t {
	case "echo":
		fmt.Println(t + " is a shell builtin")
	case "exit":
		fmt.Println(t + " is a shell builtin")
	case "type":
		fmt.Println(t + " is a shell builtin")
	default:
		fmt.Println(t + ": not found")
	}
}
func main() {
loop:
	for {
		// Uncomment this block to pass the first stage
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			break loop
		}

		command = strings.TrimSpace(command)

		switch {
		case command == "exit 0":
			os.Exit(0)
			break loop
		case strings.HasPrefix(command, "echo"):
			after, hasPrefix := strings.CutPrefix(command, "echo")
			if hasPrefix {
				fmt.Println(strings.TrimSpace(after))
			}
		case strings.HasPrefix(command, "type"):
			after, hasFound := strings.CutPrefix(command, "type")
			if hasFound {
				checkIfBuiltin(strings.TrimSpace(after))
			}
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
