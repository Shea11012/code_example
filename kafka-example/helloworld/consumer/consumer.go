package consumer

import (
	"kafka-example/conf"
	"log"

	"github.com/Shopify/sarama"
)

// 独立消费者，但服务挂掉重启后，不能从已经消费了的offset后面继续消费
func Consumer(topic string) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{conf.HOST}, config)
	if err != nil {
		log.Fatalln("NewConsumer err:", err)
	}

	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalln("ConsumerPartition err:", err)
	}

	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		log.Printf("partitionid: %d,offset: %d,key: %#v,value: %s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
	}
}
