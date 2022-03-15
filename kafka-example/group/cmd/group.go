package main

import (
	"kafka-example/conf"
	"kafka-example/group"
)

func main() {
	topic := conf.Topic
	groupName := "g1"
	go group.ConsumerGroup(topic, groupName, "c1")
	go group.ConsumerGroup(topic, groupName, "c2")
	group.ConsumerGroup(topic, groupName, "c3")
}
