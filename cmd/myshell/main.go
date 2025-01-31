package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func get_paths(path string) []string {
	if runtime.GOOS == "windows" {
		return strings.Split(path, ";")
	} else {
		return strings.Split(path, ":")
	}
}
func checkExecPath(cmd string) string {
	fullPaths := os.Getenv("PATH")
	/*
		if fullPaths != "" {
			fmt.Println("PATH is :" + fullPaths)
		}
	*/
	paths := get_paths(fullPaths)

	l := len(paths)
	if l > 0 {

		fmt.Println("Paths found: ", l)
		fmt.Println(strings.Join(paths, "'\n'"))

		for i := 0; i < len(paths); i++ {
			_, err := exec.LookPath(paths[i])
			if err == nil {
				return cmd + " is " + paths[i]
			}
		}
	}
	return cmd + ": not found"
}
func checkIfBuiltin(cmd string) {

	s := checkExecPath(cmd)

	switch cmd {
	case "echo":
		fmt.Println(cmd + " is a shell builtin")
	case "exit":
		fmt.Println(cmd + " 0 is a shell builtin")
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
		default:
			fmt.Println(command + ": command not found")
		}
		/*
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
		*/
	}
}
