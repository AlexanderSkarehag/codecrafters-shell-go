package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func checkIfExcec(s string) string {
	path, err := exec.LookPath(s)
	if err == nil {
		return path
	}
	return ""

}
func checkIfBuiltin(cmd string, execPath string) {

	s := ""
	/*
		res, err := exec.LookPath(cmd)
	*/
	if execPath == "" {
		s = cmd + ": not found"
	} else {
		s = cmd + " is " + execPath
	}

	switch cmd {
	case "echo":
		fmt.Println(cmd + " is a shell builtin")
	case "exit":
		fmt.Println(cmd + " is a shell builtin")
	case "type":
		fmt.Println(cmd + " is a shell builtin")
	default:
		fmt.Println(s)
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

		paths := strings.SplitN(command, " ", 2)

		cmd := paths[0]
		args := ""
		if len(paths) > 1 {
			args = paths[1]
		}
		execPath := checkIfExcec(cmd)

		switch cmd {
		case "exit":
			if args == "0" {
				os.Exit(0)
				break loop
			}
		case "echo":
			fmt.Println(strings.TrimSpace(args))
		case "type":
			checkIfBuiltin(strings.TrimSpace(args), execPath)
		default:
			if execPath != "" {
				exec.Command(cmd, args)
			} else {
				fmt.Println(command + ": command not found")
			}
		}
	}
}
