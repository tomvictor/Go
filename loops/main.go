package main

import "fmt"

func infinite_loop() {
	fmt.Println("infinite_loop")
	i := 99
	for {
		fmt.Println(i)
		i++
	}
}

func loop_break() {
	fmt.Println("loop_break")
	i := 99
	for {
		fmt.Println(i)
		break
	}
}

func condition_loop() {
	fmt.Println("condition_loop")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func counter_loop() {
	fmt.Println("counter_loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Go loops")

	//infinite_for_loop()
	loop_break()
	condition_loop()
	counter_loop()
}
