package main

import "fmt"

type Engine interface {
	Start()
}

type PetrolEngine struct {
	Name string
}

func (e PetrolEngine) Start() {
	fmt.Println("starting...", e.Name)
}

type GasEngine struct {
	Name string
}

func (e GasEngine) Start() {
	fmt.Println("starting...", e.Name)
}

func Start(e Engine) {
	e.Start()
}

func main() {
	fmt.Println("Interfaces")
	pEngine := PetrolEngine{
		Name: "Petrol",
	}
	gEngine := GasEngine{
		Name: "Gas",
	}

	Start(pEngine)
	Start(gEngine)
}
