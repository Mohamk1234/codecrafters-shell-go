package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

var cmds = make(map[string]func([]string))

func exitCmd(args []string) {
	if args[0] == "0" {
		os.Exit(0)
	}
}

func typeCmd(args []string) {
	cmd := args[0]
	_, ok := cmds[cmd]
	if ok {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}
	commandPath, found := findPath(cmd)
	if found {
		fmt.Printf("%s is %s\n", cmd, commandPath)
		return
	}
	fmt.Printf("%s not found\n", cmd)

}

func findPath(cmd string) (string, bool) {
	if strings.HasPrefix(cmd, "/") {
		return cmd, true
	}
	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, p := range paths {
		filePath := path.Join(p, cmd)
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			return filePath, true
		}
	}
	return "", false
}

func echoCmd(args []string) {
	res := ""
	for _, a := range args {
		res += " " + a
	}
	fmt.Print(res + "\n")
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")
	cmds["exit"] = exitCmd
	cmds["type"] = typeCmd
	cmds["echo"] = echoCmd
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		trimmedCommand := strings.Split(strings.TrimSpace(command), " ")

		cmd, ok := cmds[trimmedCommand[0]]

		if ok {
			cmd(trimmedCommand[1:])
			continue
		}
		path, ok := findPath(trimmedCommand[0])
		if ok {
			out, err := exec.Command(path, trimmedCommand[1:]...).Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(string(out) + "\n")
			continue
		}

		fmt.Print(trimmedCommand[0] + ": command not found\n")

	}

}
