package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var PATH string = os.Getenv("PATH")

func checkIfExecutable(command string) (bool, string) {
	paths := strings.Split(PATH, ":")
	for _, dir := range paths {
		filePath := fmt.Sprintf("%s/%s", dir, command)
		if info, err := os.Stat(filePath); err == nil {

			const (
				ExecAny os.FileMode = 0111 // Any execute permission
			)

			mode := info.Mode().Perm()

			if mode&ExecAny != 0 {
				return true, filePath
			}
		}
	}
	return false, ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := []string{"echo", "exit", "type"}

inputLoop:
	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		fields := strings.Fields(input)
		if len(fields) == 0 {
			continue
		}
		command := fields[0]

		switch command {
		case "echo":
			fmt.Printf("%s\n", strings.Join(fields[1:], " "))
		case "exit":
			break inputLoop
		case "type":
			test := strings.Join(fields[1:], " ")
			if slices.Contains(commands, test) {
				fmt.Printf("%s is a shell builtin\n", test)
			} else {
				if found, path := checkIfExecutable(test); found {
					fmt.Printf("%s is %s\n", test, path)
				} else {
					fmt.Printf("%s: not found\n", test)
				}
			}
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}

}
