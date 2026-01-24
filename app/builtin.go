package main

import (
	"fmt"
	"log"
	"os"
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

func executePwd() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentDir)
}

func executeCd(args []string) {
	if len(args) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		err = os.Chdir(homeDir)
		if err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", homeDir)
		}
		return
	}

	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", dir)
	}
}
