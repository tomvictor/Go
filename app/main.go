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
			err := menu.AddItem()
			if err != nil {
				fmt.Println(fmt.Errorf("Error Occured: %w", err))
			}
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}

}
