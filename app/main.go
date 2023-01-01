package main

import (
	"app/menu"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {

loop:

	for {
		menu.PrintInfo()

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			menu.AddItem()
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}

}
