package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func main() {
	producer, err := sarama.NewAsyncProducer([]string{"10.3.10.32:9091"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, errors int
ProducerLoop:
	for {
		select {
		case producer.Input() <- &sarama.ProducerMessage{Topic: "dataman_test", Key: nil, Value: sarama.StringEncoder("testing 123")}:
			enqueued++
		case err = <-producer.Errors():
			log.Println("Failed to produce message", err)
			errors++
		case <-signals:
			break ProducerLoop
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}
