package main

import (
	"kafka-example/conf"
	"kafka-example/helloworld/consumer"
)

func main() {
	topic := conf.Topic
	consumer.Consumer(topic)
}
