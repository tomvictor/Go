package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {

	msgCh := make(chan string)

	go FetchFromCli(msgCh)
	go PrintOutCli(msgCh)

	time.Sleep(time.Minute)

}

func PrintOutCli(msgCh chan string) {
	var data string
	for {
		select {
		case data = <-msgCh:
			fmt.Println(data)
		}
	}
}

func FetchFromCli(msg chan string) {
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
		msg <- string(buff)
	}
}
