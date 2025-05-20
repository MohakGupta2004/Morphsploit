package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MohakGupta2004/Morphsploit/commands"
)

func startShell() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("morphsploit~# ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}
		command := args[0]
		commandArgs := args[1:]
		switch command {
		case "echo":
			commands.Echo(commandArgs)
		case "exit":
			commands.ExitSploit()
		}

	}
}
