package main

import (
	"fmt"
	"strings"

	"github.com/MohakGupta2004/Morphsploit/commands"
	"github.com/chzyer/readline"
)

type UseMode struct {
	auth   bool
	module string
}

func startShell() {
	rl, err := readline.New("morphsploit~# ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	history := []string{}
	useMode := UseMode{
		auth:   false,
		module: "",
	}
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
			if useMode.auth == true {
				useMode.auth = false
				rl.SetPrompt("morphsploit~# ")
				break
			} else {
				commands.ExitSploit()
				return
			}
		case "use":
			err, module := commands.Use(commandArgs)
			if err != nil {
				fmt.Println(err)
				break
			}
			useMode.auth = true
			useMode.module = module
			shell := fmt.Sprintf("morphsploit (%s)~# ", strings.Split(commandArgs[0], "/")[2])
			rl.SetPrompt(shell)
		case "show":
			if useMode.auth == false {
				fmt.Println("Module not specified")
				break
			}
			if len(commandArgs) == 0 {
				fmt.Println("Usage: show options")
				break
			}
			if commandArgs[0] != "options" {
				break
			}
			commands.Opt(useMode.module)
		case "set":

		default:
			fmt.Println("Unknown Command")
		}
	}
}
