package mq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

func TestConsumer(topic string, channel string, hdlr nsq.HandlerFunc) error {
	config := nsq.NewConfig()
	config.MaxInFlight = 10
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return err
	}

	consumer.AddConcurrentHandlers(hdlr, 1)
	if err := consumer.ConnectToNSQLookupd("192.168.2.109:4161"); err != nil {
		return err
	}
	<-consumer.StopChan
	return nil
}

func OnMessage(message *nsq.Message) error {
	fmt.Printf(string(message.Body))
	return nil
}
