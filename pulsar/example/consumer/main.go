package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

var topic = "test"
var URL = "pulsar://localhost:6650"

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               URL,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: "test_sub_1",
		Type:             pulsar.Shared,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()
	ctx := context.Background()
	for {
		msg, err := consumer.Receive(ctx)
		if err != nil {
			log.Println(err)
			break
		}

		payload := string(msg.Payload())
		fmt.Printf("Received message msgId: %#v -- content: '%s'\n", msg.ID().EntryID(), payload)
		err = consumer.Ack(msg)
		if err != nil {
			log.Println(err)
			break
		}

		time.Sleep(time.Second * 3)
	}

	fmt.Println("消费者停止")
}
