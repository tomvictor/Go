package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cli := exec.Command("./sensor/sensor")
	pipe, err := cli.StdoutPipe()
	if err != nil {
		panic(err)
	}
	err = cli.Start()
	if err != nil {
		panic(err)
	}

	buff := make([]byte, 20)
	for {
		pipe.Read(buff)
		fmt.Println(string(buff))
	}
}
