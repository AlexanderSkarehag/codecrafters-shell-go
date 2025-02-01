package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func checkIfExcec(s string) bool {
	_, err := exec.LookPath(s)
	return err == nil

}
func checkIfBuiltin(cmd string) {

	builtin := []string{"echo", "exit", "type", "pwd", "cd"}
	s := ""
	path, err := exec.LookPath(cmd)
	if err != nil {
		s = cmd + ": not found"
	} else {
		s = cmd + " is " + path
	}

	if slices.Contains(builtin, cmd) {
		printShellBuiltin(cmd)
	} else {
		fmt.Println(s)
	}
}
func printShellBuiltin(s string) {
	fmt.Println(s + " is a shell builtin")
}
func getDirectoryPath(s string) string {
	if s == "~" {
		p := os.Getenv("HOME")
		if p == "" {
			p2, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error in ~ directorypath")
			}
			p = p2
		}

		return strings.TrimSpace(p)
	}

	p, err := filepath.Abs(s)
	if err != nil {
		fmt.Println("Error1!")
	}
	return p
}
func echo(s string) {
	s = strings.Trim(s, "'")
	fmt.Println(strings.TrimSpace(s))
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
			echo(args)
		case "type":
			checkIfBuiltin(strings.TrimSpace(args))
		case "pwd":
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error!")
			}
			fmt.Println(pwd)
		case "cd":
			p := getDirectoryPath(args)
			if err := os.Chdir(p); err != nil {
				fmt.Println("cd: " + args + ": No such file or directory")
			}
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
