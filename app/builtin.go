package main

import (
	"fmt"
	"slices"
	"strings"
)

func executeEcho(args []string) {
	output := ""
	for i, arg := range args {
		output += arg
		if i < len(args)-1 {
			output += " "
		}
	}
	fmt.Println(output)
}

func executeType(args []string) {
	if len(args) == 0 {
		fmt.Println("type: missing argument")
		return
	}

	command := strings.Join(args, " ")
	if slices.Contains(builtinCommands, command) {
		fmt.Printf("%s is a shell builtin\n", command)
	} else {
		if found, path := checkIfExecutable(command); found {
			fmt.Printf("%s is %s\n", command, path)
		} else {
			fmt.Printf("%s: not found\n", command)
		}
	}
}
