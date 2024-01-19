package main

import (
	"bufio"
	"fmt"
	"os"
)

var G = new(Game)
var P = new(Player)

func main() {
	initGame()

	var command string
	input := bufio.NewReader(os.Stdin)
	for {
		command, _ = input.ReadString('\n')
		fmt.Printf("command is %s\n", command)
		fmt.Println(handleCommand(command))
	}
}