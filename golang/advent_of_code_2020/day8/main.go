package main

import (
	"fmt"
	"strconv"
	"strings"

	"../util"
)

type command struct {
	name   string
	amount int
}

// 236 jmp -125
const del = "\n"
const accMark = 5
const lastLen = 5

func main() {
	text := util.ReadFileToString("./input.txt")
	lines := strings.Split(text, del)

	commands := parseInputToCommands(lines)

	// result := runCommands(commands)

	// fmt.Println()
	// fmt.Println("RESULT: ", result)

	for i, cmd := range commands {
		var cmdsCopy []command = make([]command, len(commands))
		copy(cmdsCopy, commands)

		if cmd.name == "jmp" {
			cmdsCopy[i].name = "nop"
		} else if cmd.name == "nop" {
			cmdsCopy[i].name = "jmp"
		} else {
			continue
		}
		bruteResult := runCommands(cmdsCopy)
		if bruteResult {
			fmt.Println("FOUND", cmd)
		}
	}

}

func runCommands(cmds []command) bool {
	acc := 0
	j := 0
	mark := make([]bool, len(cmds))
	for i := 0; i < len(cmds); i++ {
		var cmd command = cmds[i]

		// Our exit case (add `|| mark[i]` for test 1)
		j++
		if j > 100000 {
			// fmt.Println("J:", j)
			return false
		}

		if cmd.name == "acc" {
			acc += cmd.amount
		} else if cmd.name == "jmp" {
			i += (cmd.amount - 1)
		}

		if i < 0 {
			return false
		}
		mark[i] = true
	}
	fmt.Println("FIN", acc)
	return true
}

func parseInputToCommands(lines []string) []command {
	commands := []command{}

	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		name := lineSplit[0]
		amount := lineSplit[1]
		if string(amount[0]) == "+" {
			amount = amount[1:]
		}
		convertedAmount, _ := strconv.Atoi(amount)

		var cmd command = command{name: name, amount: convertedAmount}
		commands = append(commands, cmd)
	}

	return commands
}
