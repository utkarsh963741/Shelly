package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var PATH string = os.Getenv("PATH")

func executeExternal(fields []string) {
	command := fields[0]
	args := fields[1:]

	if found, _ := checkIfExecutable(command); found {
		executeExternalCommand(command, args)
	} else {
		fmt.Printf("%s: command not found\n", command)
	}
}

func checkIfExecutable(command string) (bool, string) {
	currOS := runtime.GOOS
	delimiter := string(os.PathSeparator)
	seperator := string(os.PathListSeparator)
	extension := ""

	switch currOS {
	case "windows":
		delimiter = "\\"
		seperator = ";"
		extension = ".exe"
	case "linux":
		delimiter = "/"
		seperator = ":"
		extension = ""
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
				const ExecAny os.FileMode = 0111
				mode := info.Mode().Perm()
				if mode&ExecAny != 0 {
					return true, filePath
				}
			}
		}
	}
	return false, ""
}

func executeExternalCommand(command string, args []string) {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("%s failed with %s\n", command, err)
	}
}
