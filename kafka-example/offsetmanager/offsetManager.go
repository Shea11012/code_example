package offsetmanager

import (
	"kafka-example/conf"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

/*
offsetManager 记录下每次消费的位置，避免 offsetOldest 重复消费或者 offsetNewest 漏掉部分消息
*/

func OffsetManager(topic string) {
	config := sarama.NewConfig()
	// 开启自动提交
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	client, err := sarama.NewClient([]string{conf.HOST}, config)
	if err != nil {
		log.Fatalln("NewClient err:", err)
	}

	defer client.Close()

	// 根据groupID 来区分不同的 consumer，每次提交的offset信息都是与group相关联
	offsetManager, err := sarama.NewOffsetManagerFromClient("group1", client)
	if err != nil {
		log.Fatalln("NewOffsetManagerFromClient err:", err)
	}
	defer offsetManager.Close()

	partitionOffsetManager, err := offsetManager.ManagePartition(topic, 0)
	if err != nil {
		log.Fatalln("ManagePartition err:", err)
	}
	defer partitionOffsetManager.Close()

	defer offsetManager.Commit()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatalln("NewConsumerFromClient err:", err)
	}

	nextOffset, _ := partitionOffsetManager.NextOffset()
	println("nextOffset:", nextOffset)
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, nextOffset)
	if err != nil {
		log.Fatalln("ConsumePartition err:", err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		log.Printf("partitionID: %d,offset: %d,value: %s\n", msg.Partition, msg.Offset, msg.Value)

		// MarkOffset 标记最后一个消费，此时只会存在内存中，需要 commit 之后才会提交到 kafka
		partitionOffsetManager.MarkOffset(msg.Offset+1, "modified metadata")
	}

}
