package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var builtinCommands = []string{"echo", "type", "pwd", "exit"}

func runShell() {
	reader := bufio.NewReader(os.Stdin)

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
			executeEcho(fields[1:])
		case "type":
			executeType(fields[1:])
		case "pwd":
			executePwd()
		case "exit":
			break inputLoop
		default:
			executeExternal(fields)
		}
	}
}
