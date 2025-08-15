package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/ummuys/level_0/cmd/producer/jsons"
)

func main() {

	err := godotenv.Load(".env.producer")
	if err != nil {
		log.Fatal("no env file: ", err)
	}

	topic := os.Getenv("KAFKA_TOPIC")
	broker := os.Getenv("KAFKA_BROKER")
	transactionalId := os.Getenv("KAFKA_TID")

	if broker == "" || topic == "" || transactionalId == "" {
		log.Fatal("kafka: KAFKA_BROKERS, KAFKA_TOPIC, KAFKA_TID are required")
	}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(broker),
		kgo.TransactionalID(transactionalId),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer cl.Close()

	send(cl, topic, jsons.GoodJSONs)
	send(cl, topic, jsons.BadJSONs)

}

func send(cl *kgo.Client, topic string, orders []string) {
	for _, ord := range orders {

		if err := cl.BeginTransaction(); err != nil {
			log.Fatal("can't begin the thansaction: ", err)
		}

		rec := &kgo.Record{
			Topic: topic,
			Key:   []byte("order-key"),
			Value: []byte(ord),
			Headers: []kgo.RecordHeader{
				{Key: "content-type", Value: []byte("application/json")},
				{Key: "event-type", Value: []byte("order.created")},
				{Key: "schema-version", Value: []byte("v1")},
			},
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		res := cl.ProduceSync(ctx, rec)
		cancel()

		if err := res.FirstErr(); err != nil {
			_ = cl.EndTransaction(context.Background(), false) // abort
			log.Fatal("produce:", err)
		}
		if err := cl.EndTransaction(context.Background(), true); err != nil {
			log.Fatal("commit:", err)
		}
		log.Println("Tx committed")

	}
}
