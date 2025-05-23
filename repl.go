package main

import (
	"fmt"
	"strings"

	"github.com/MohakGupta2004/Morphsploit/commands"
	"github.com/chzyer/readline"
)

func startShell() {
	rl, err := readline.New("morphsploit~# ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	history := []string{}
	useMode := false
	for {
		line, err := rl.Readline()
		if err != nil { // Handle Ctrl+C or EOF
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		history = append(history, line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}
		command := args[0]
		commandArgs := args[1:]
		switch command {
		case "echo":
			commands.Echo(commandArgs)
		case "exit":
			if useMode == true {
				useMode = false
				rl.SetPrompt("morphsploit~# ")
				break
			} else {
				commands.ExitSploit()
				return
			}
		case "use":
			err := commands.Use(commandArgs)
			if err != nil {
				fmt.Println(err)
				break
			}
			useMode = true
			shell := fmt.Sprintf("morphsploit (%s)~# ", strings.Split(commandArgs[0], "/")[2])
			rl.SetPrompt(shell)
		case "show":
			if commandArgs[0] != "options" {
				break
			}
			if useMode == false {
				fmt.Println("Module not specified")
				break
			}
		default:
			fmt.Println("Unknown Command")
		}
	}
}
