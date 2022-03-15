package producer

import (
	"context"
	"kafka-example/conf"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

func connectBroker() *sarama.Broker {
	broker := sarama.NewBroker(conf.HOST)
	config := sarama.NewConfig()
	config.Version = sarama.MaxVersion
	err := broker.Open(config)
	if err != nil {
		log.Fatalln("Open err:", err)
	}

	connected, err := broker.Connected()
	if err != nil {
		log.Fatalln("connected err:", err)
	}

	log.Println(connected)

	return broker
}

func CreateTopic(topic string) {
	broker := connectBroker()
	defer broker.Close()

	topicDetails := map[string]*sarama.TopicDetail{}
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}

	topicDetails[topic] = topicDetail

	topicRequest := &sarama.CreateTopicsRequest{
		Timeout:      time.Second * 10,
		TopicDetails: topicDetails,
	}

	topicResponse, err := broker.CreateTopics(topicRequest)
	if err != nil {
		log.Fatalln("CreateTopics err:", err)
	}

	for _, v := range topicResponse.TopicErrors {
		log.Printf("error key: %d, error msg: %s\n", v.Err, v.Error())
	}
}

// 同步发送
func Produce(topic string, limit int) {
	config := sarama.NewConfig()
	// 同步生产者必须将 success 和 errors 都开启
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{conf.HOST}, config)
	if err != nil {
		log.Fatalln("NewSyncProducer err:", err)
	}

	defer producer.Close()

	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   nil,
			Value: sarama.StringEncoder(str),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatalln("SendMessage err:", err)
			return
		}

		log.Printf("Producer partitionId: %d,offset: %d,value: %s\n", partition, offset, str)
	}
}

// 异步发送
func AsyncProducer(topic string, limit int) {
	config := sarama.NewConfig()
	// 异步生产者，一般只需要开启 errors 就可以
	config.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer([]string{conf.HOST}, config)
	if err != nil {
		log.Fatalln("NewAsyncProducer err:", err)
	}

	var wg sync.WaitGroup

	// 异步生产者必须将返回值从 errors 或 success 中读出，不然发生阻塞将只能发送出去一条消息
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range producer.Successes() {
			log.Printf("producer success: key: %#v msg: %s\n", msg.Key, msg.Value)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for e := range producer.Errors() {
			log.Printf("producer error: err: %v,msg: %+v\n", e.Err, e.Msg)
		}
	}()

	for i := 0; i < limit; i++ {
		str := strconv.Itoa(int(time.Now().UnixNano()))
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Key:   nil,
			Value: sarama.StringEncoder(str),
		}

		// 异步发送只是写入了内存，使用超时控制防止 input <- msg 阻塞
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*10)
		select {
		case producer.Input() <- msg:
		case <-ctx.Done():
		}

		cancel()
	}

	producer.AsyncClose()
	wg.Wait()
}
