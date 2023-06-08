package hello

import (
	"github.com/go-gorf/gorf"
)

func setup() error {
	// Add setup here
	mqttChannel := make(chan string)
	mqttClient := MqRunner(mqttChannel)
	go mqttPublisher(mqttClient)
	go mqttPrintMsg(mqttChannel)
	return nil
}

var HelloApp = gorf.GorfBaseApp{
	Name:         "Hello",
	RouteHandler: Urls,
	SetUpHandler: setup,
}
