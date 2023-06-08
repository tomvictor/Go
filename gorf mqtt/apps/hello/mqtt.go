package hello

import (
	"fmt"
	"os"
	"time"

	// MQTT import the Paho Go MQTT library
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const MQTT_PUB_TOPIC = "gorf/mqtt"

func MqRunner(mqttChannel chan string) MQTT.Client {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt.eclipseprojects.io:1883")
	opts.SetClientID("go-simple")
	//opts.SetDefaultPublishHandler()

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	token := c.Subscribe(MQTT_PUB_TOPIC, 0, func(client MQTT.Client, msg MQTT.Message) {
		data := fmt.Sprintf("(%s): %s \n", msg.Topic(), msg.Payload())
		mqttChannel <- data
	})
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	return c
}

func mqttPrintMsg(mqttChannel chan string) {
	for {
		select {
		case msg := <-mqttChannel:
			fmt.Println(msg)
		}
	}
}

func mqttPublisher(c MQTT.Client) {
	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for i := 0; i < 1000; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish(MQTT_PUB_TOPIC, 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe(MQTT_PUB_TOPIC); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}
