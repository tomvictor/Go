package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

type menu []menuItem

var in = bufio.NewReader(os.Stdin)

// only functions with Capfirst will be publicaly available
func (m menu) Print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for key, value := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", key, value)
		}
	}
}

func (m *menu) AddItem() error {
	name, _ := in.ReadString('\n')
	for _, item := range data {
		if item.name == name {
			return errors.New("Item already exists")
		}
	}
	data = append(*m, menuItem{name: name, prices: make(map[string]float64)})
	return nil
}

func AddItem() error {
	return data.AddItem()
}

func Print() {
	data.Print()
}

func PrintInfo() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	fmt.Println("2) Add item")
	fmt.Println("q) quit")
}
