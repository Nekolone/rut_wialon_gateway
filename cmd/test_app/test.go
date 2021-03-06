package main

import (
	"log"
	"time"
)

func main() {
	log.Print(time.Now().Format("2006-01-02 15:04:05.00"))

}

//import the Paho Go MQTT library

// define a function for the default message handler
// var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
// 	fmt.Printf("TOPIC: %s\n", msg.Topic())
// 	fmt.Printf("MSG: %s\n", msg.Payload())
// }

// var msgHandl MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
// 	fmt.Printf("get TOPIC: %s\n", msg.Topic())
// 	fmt.Printf("get MSG: %s\n", msg.Payload())
// }

// func main() {
// 	// create a ClientOptions struct setting the broker address, clientid, turn
// 	// off trace output and set the default message handler
// 	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.35.63:18883")
// 	// opts.SetClientID("go-simple")
// 	// opts.SetDefaultPublishHandler(f)

// 	// create and start a client using the above ClientOptions
// 	c := MQTT.NewClient(opts)
// 	if token := c.Connect(); token.Wait() && token.Error() != nil {
// 		panic(token.Error())
// 	}
// 	c.IsConnected()

// 	// subscribe to the topic /go-mqtt/sample and request messages to be delivered
// 	// at a maximum qos of zero, wait for the receipt to confirm the subscription
// 	if token := c.Subscribe("#", 0, msgHandl); token.Wait() && token.Error() != nil {
// 		fmt.Println(token.Error())
// 		os.Exit(1)
// 	}

// 	// c.SubscribeMultiple(map[string]byte{"modbusData": 0, "modbusData2": 0}, msgHandl)

// 	// Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
// 	// from the server after sending each message
// 	// go publishThis(c)

// 	log.Println("stopped")
// 	time.Sleep(10 * time.Second)

// 	// unsubscribe from /go-mqtt/sample
// 	if token := c.Unsubscribe("#"); token.Wait() && token.Error() != nil {
// 		fmt.Println(token.Error())
// 		os.Exit(1)
// 	}

// 	c.Disconnect(250)
// }

// func publishThis(c MQTT.Client) {
// 	for i := 0; i < 5; i++ {
// 		text := fmt.Sprintf("this is msg #%d!", i)
// 		token := c.Publish("go-mqtt/sample", 0, false, text)
// 		token.Wait()
// 	}
// }
