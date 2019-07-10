package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Add variable
	var lokern [5]int
	lokern[0] = 1
	lokern[1] = 2
	lokern[2] = 3
	lokern[3] = 4
	lokern[4] = 5

	/** Print masing -masing value dari variable diatas **/

	loker := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := loker.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "init":
	case "status":
	case "input":
	case "leave":
	case "find":
	case "search":
	case "exit":
		os.Exit(0)

	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
