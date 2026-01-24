package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
)

var PATH string = os.Getenv("PATH")

func checkIfExecutable(command string) (bool, string) {
	currOS := runtime.GOOS
	delimiter := string(os.PathSeparator)
	seperator := string(os.PathListSeparator)
	extension := ""

	switch currOS {
	case "windows":
		{
			// fmt.Println("Running on Windows")
			delimiter = "\\"
			seperator = ";"
			extension = ".exe"
		}

	case "linux":
		{
			// fmt.Println("Running on Linux")
			delimiter = "/"
			seperator = ":"
			extension = ""
		}

	default:
		fmt.Printf("Running on %s\n", currOS)

	}

	paths := strings.Split(PATH, seperator)
	for _, dir := range paths {
		filePath := fmt.Sprintf("%s%s%s%s", dir, delimiter, command, extension)
		if info, err := os.Stat(filePath); err == nil {

			if currOS == "windows" {
				return true, filePath
			} else {
				const (
					ExecAny os.FileMode = 0111 // Any execute permission
				)
				mode := info.Mode().Perm()
				if mode&ExecAny != 0 {
					return true, filePath
				}
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
			if found, _ := checkIfExecutable(command); found {
				args := fields[1:]
				cmd := exec.Command(command, args...)

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				err := cmd.Run()
				if err != nil {
					log.Fatalf("%s failed with %s\n", command, err)
				}
			} else {
				fmt.Printf("%s: command not found\n", command)
			}
		}
	}

}
