package main

import (
	"kafka-example/conf"
	"kafka-example/helloworld/producer"
)

func main() {
	topic := conf.Topic
	producer.CreateTopic(topic)
	producer.Produce(topic, 10)
	producer.AsyncProducer(topic, 10)
}
