package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

var topic = "test"
var URL = "pulsar://localhost:6650"

type Message struct {
	Id   int64  `json:"id"`
	Time string `json:"time"`
	Data string `json:"data"`
}

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

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer producer.Close()
	ctx := context.Background()
	count := 1
	for i := 0; i < count; i++ {
		message := &Message{
			Id:   int64(i),
			Time: time.Now().Format(time.DateTime),
			Data: fmt.Sprintf("{\"msg\":\"消息-%d\"}", i),
		}

		b, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			break
		}

		_, err = producer.Send(ctx, &pulsar.ProducerMessage{
			Payload: b,
		})

		if err != nil {
			fmt.Println("Failed to publish message", err)
			break
		}

		fmt.Printf("%d 发送成功\n", i)
	}

	print("生产者发送完毕")
}
