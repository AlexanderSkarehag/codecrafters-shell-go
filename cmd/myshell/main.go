package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func checkIfExcec(s string) bool {
	_, err := exec.LookPath(s)
	return err == nil

}
func checkIfBuiltin(cmd string) {

	s := ""
	path, err := exec.LookPath(cmd)
	if err != nil {
		s = cmd + ": not found"
	} else {
		s = cmd + " is " + path
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

	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")
loop:
	for {
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
		isExec := checkIfExcec(cmd)

		switch cmd {
		case "exit":
			if args == "0" {
				os.Exit(0)
				break loop
			}
		case "echo":
			fmt.Println(strings.TrimSpace(args))
		case "type":
			checkIfBuiltin(strings.TrimSpace(args))
		case "pwd":
			ex, err := os.Executable()
			if err != nil {
				fmt.Println("Error!")
			}
			fmt.Println(filepath.Dir(ex))
		default:
			if isExec {
				c := exec.Command(cmd, strings.TrimSpace(args))
				output, err := c.Output()
				if err != nil {
					log.Printf("Command finished with error: %v", err)
				}
				fmt.Println(strings.TrimSpace(string(output)))

			} else {
				fmt.Println(command + ": command not found")
			}
		}
		fmt.Fprint(os.Stdout, "$ ")
	}
}
