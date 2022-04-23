package core

import (
	"fmt"
	"github.com/ZongweiBai/golang-in-action/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	config.LOG.Infof("Received message: %s from topic: %s", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	config.LOG.Infof("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	config.LOG.Errorf("Connect lost: %v", err)
}

func InitMqttClient() {
	if config.CONFIG.MQTT.Enabled == false {
		return
	}
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.CONFIG.MQTT.Broker, config.CONFIG.MQTT.Port))
	opts.SetClientID(config.CONFIG.MQTT.ClientId)
	opts.SetUsername(config.CONFIG.MQTT.UserName)
	opts.SetPassword(config.CONFIG.MQTT.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(100)

	sub(client)
	publish(client)
}

func publish(client mqtt.Client) {
	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish("topic/test", 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}

func sub(client mqtt.Client) {
	topic := config.CONFIG.MQTT.Topics
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	config.LOG.Infof("Subscribed to topic: %s", topic)
}
