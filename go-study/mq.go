package util

import (
	"fmt"
	"sync"

	"github.com/Dataman-Cloud/omega-metrics/config"
	log "github.com/Sirupsen/logrus"
	"github.com/streadway/amqp"
)

const (
	Metrics_exchange       string = "cluster_info"
	Master_metrics_routing string = "master_metrics"
	Master_state_routing   string = "master_state"
	Slave_metrics_routing  string = "slave_metrics"
	Slave_state_routing    string = "slave_state"
)

func failOnError(err error, msg string) error {
	if err != nil {
		log.Errorf("%s: %s", msg, err)
	}
	return err
}

var mq *amqp.Connection

func MQ() *amqp.Connection {
	if mq != nil {
		return mq
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	InitMQ()
	defer mutex.Unlock()

	return mq
}

func Publish(name string, message []byte) error {
	mq := MQ()
	channel, err := mq.Channel()
	if err != nil {
		log.Error("can't get channel", err)
		return err
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(name, "fanout", true, false, false, false, nil)
	if err != nil {
		log.Error("can't declare exchange", err)
		return err
	}

	err = channel.Publish(
		name,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		log.Error("can't publish message", err)
		return err
	}
	return nil
}

// func Subscribe(name string, handler func([]byte)) error {
// 	mq := MQ()
// 	channel, err := mq.Channel()
// 	if err != nil {
// 		log.Error("can't get channel", err)
// 		return err
// 	}

// 	err = channel.ExchangeDeclare(name, "fanout", true, false, false, false, nil)
// 	if err != nil {
// 		log.Error("can't declare exchange", err)
// 		return err
// 	}

// 	queue, err := channel.QueueDeclare("", false, false, true, false, nil)
// 	if err != nil {
// 		log.Error("can't declare queue", err)
// 		return err
// 	}

// 	err = channel.QueueBind(queue.Name, "", name, false, nil)
// 	if err != nil {
// 		log.Error("can't bind queue ", err)
// 		return err
// 	}

// 	messages, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
// 	if err != nil {
// 		log.Error("can't consume ", err)
// 		return err
// 	}

// 	go func() {
// 		defer channel.Close()
// 		for message := range messages {
// 			handler(message.Body)
// 		}
// 	}()

// 	return nil
// }

func MetricsPublish(exchange string, message []byte) error {
	mq := MQ()
	channel, err := mq.Channel()
	failOnError(err, "can't get channel")
	defer channel.Close()

	err = channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	failOnError(err, "can't declare exchange")

	err = channel.Publish(
		exchange,
		Master_metrics_routing,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	failOnError(err, "can't pulish message")
	return nil
}

func MetricsSubscribe(exchange string, routingkey string, handler func(string, []byte)) error {

	mq := MQ()
	channel, err := mq.Channel()
	failOnError(err, "can't get channel")

	err = channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil)
	failOnError(err, "can't declare exchange")

	err = channel.QueueBind(routingkey, routingkey, exchange, false, nil)
	failOnError(err, "can't bind queue")

	messages, err := channel.Consume(routingkey, "", true, false, false, false, nil)
	failOnError(err, "can't consume")

	go func() {
		defer channel.Close()
		for message := range messages {
			handler(message.RoutingKey, message.Body)
		}
	}()

	return nil
}

func InitMQ() {
	conf := config.Pairs()
	opts := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		conf.Mq.User, conf.Mq.Password, conf.Mq.Host, conf.Mq.Port)
	var err error
	mq, err = amqp.Dial(opts)
	if err != nil {
		log.Error("got err", err)
		log.Fatal("can't dial mq server: ", opts)
		panic(-1)
	}
	log.Debug("initialized MQ")
}

func DestroyMQ() {
	log.Info("destroying MQ")
	if mq != nil {
		mq.Close()
	}
}
