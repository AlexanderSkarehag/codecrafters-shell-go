package main

import (
	"bufio"
	"fmt"
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
func printCommandWithoutBuiltin(s string) {
	path, err := exec.LookPath(s)
	if err != nil {
		fmt.Println(s + ": not found")
	} else {
		fmt.Println(s + " is " + path)
	}
}
func handleTypeCommand(cmd string) {
	builtin := []string{"echo", "exit", "type", "pwd", "cd"}

	if slices.Contains(builtin, cmd) {
		printShellBuiltin(cmd)
	} else {
		printCommandWithoutBuiltin(cmd)
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
func getDirectoryPaths(args []string) []string {
	l := []string{}
	for i := 0; i < len(args); i++ {
		path := getDirectoryPath(args[i])
		l = append(l, path)
	}
	return l
}
func echo(s string) {
	hasSingleQuotes := strings.HasPrefix(s, "'")
	hasDoubleQuotes := strings.HasPrefix(s, "\"")

	if hasSingleQuotes {
		l := getArgs(s, "'")
		fmt.Println(strings.Join(l, ""))
	} else if hasDoubleQuotes {
		l := getArgsWithoutSpaces(s, "\"")
		fmt.Println(strings.Join(l, ""))
	} else {
		fmt.Println(strings.Join(strings.Fields(s), " "))
	}

}
func handlePwd() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error with Getwd!")
	}
	fmt.Println(pwd)
}
func executeCommands(s string, args ...string) {
	c := exec.Command(s, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	c.Run()
}
func getArgs(s string, delimiter string) []string {
	args := strings.Split(s, delimiter)

	return args
}
func getArgsWithoutSpaces(s string, delimiter string) []string {
	list := strings.Split(s, delimiter)
	args := []string{}
	for i := 0; i < len(list); i++ {
		v := list[i]
		vList := strings.Split(v, " ")
		if len(vList) > 1 || (vList[0] != "" && vList[0] != " ") {
			args = append(args, v)
		}
	}
	return args
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
			handleTypeCommand(strings.TrimSpace(args))
		case "pwd":
			handlePwd()
		case "cd":
			p := getDirectoryPath(args)
			if err := os.Chdir(p); err != nil {
				fmt.Println("cd: " + args + ": No such file or directory")
			}
		case "cat":
			singleQ := strings.HasPrefix(args, "'")
			delimiter := ""
			if singleQ {
				delimiter = "'"
			} else {
				delimiter = "\""
			}
			l := getDirectoryPaths(getArgsWithoutSpaces(args, delimiter))
			executeCommands(cmd, l...)
		default:
			if isExec {
				executeCommands(cmd, strings.TrimSpace(args))
			} else {
				fmt.Println(command + ": command not found")
			}
		}
		fmt.Fprint(os.Stdout, "$ ")
	}
}
