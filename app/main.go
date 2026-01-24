package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

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
				fmt.Printf("%s: not found\n", test)
			}
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}

}
