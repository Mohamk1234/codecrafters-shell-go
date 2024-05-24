package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		trimmedCommand := strings.Split(strings.TrimSpace(command), " ")

		//checking if command exists
		switch trimmedCommand[0] {
		case "exit":
			if trimmedCommand[1] == "0" {
				os.Exit(0)
			}
		case "echo":
			res := ""
			for i := 1; i < len(trimmedCommand); i++ {
				res += " " + trimmedCommand[i]
			}
			fmt.Print(res + "\n")

		default:
			fmt.Print(trimmedCommand[0] + ": command not found\n")
		}
	}

}
