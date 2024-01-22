package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ecoarchie/go-mud/internal/app"
)

func main() {
	app.InitGame()

	var command string
	input := bufio.NewReader(os.Stdin)
	for {
		command, _ = input.ReadString('\n')
		fmt.Println(app.HandleCommand(command))
	}
}
