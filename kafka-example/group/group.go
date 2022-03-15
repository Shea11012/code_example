package group

import (
	"context"
	"fmt"
	"log"
	"sync"

	"kafka-example/conf"

	"github.com/Shopify/sarama"
)

// 需要实现 ConsumerGroup 接口才能作为消费者组
type ConsumerGroupHandler struct {
	name  string
	count int64
}

// 获得新 session 后的第一步，在 ConsumerClaim 之前
func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// 当所有 ConsumerClaim goroutine 都退出时
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("name: %s topic: %q partition: %d offset: %d\n", h.name, msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
		h.count++
	}

	return nil
}

func ConsumerGroup(topic, group, name string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	consumerGroup, err := sarama.NewConsumerGroup([]string{conf.HOST}, group, config)
	if err != nil {
		log.Fatalln("NewConsumerGroup err:", err)
	}
	defer consumerGroup.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler := ConsumerGroupHandler{
			name: name,
		}
		fmt.Println("running:", name)
		for {
			err = consumerGroup.Consume(ctx, []string{topic}, handler)
			if err != nil {
				log.Println("Consume err:", err)
			}

			if ctx.Err() != nil {
				return
			}
		}
	}()

	wg.Wait()
}
