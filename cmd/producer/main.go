package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/twmb/franz-go/pkg/kgo"
)

const orderJSON = `{
   "order_uid": "adsfasasdffffasdfadsfff",
   "track_number": "WBILMTESTTRACK",
   "entry": "WBIL",
   "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
   },
   "payment": {
      "transaction": "b563feb7b2b84b6test",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
   },
   "items": [
      {
         "chrt_id": 9934930,
         "track_number": "WBILMTESTTRACK",
         "price": 453,
         "rid": "ab4219087a764ae0btest",
         "name": "Mascaras",
         "sale": 30,
         "size": "0",
         "total_price": 317,
         "nm_id": 2389212,
         "brand": "Vivienne Sabo",
         "status": 202
      }
   ],
   "locale": "en",
   "internal_signature": "",
   "customer_id": "test",
   "delivery_service": "meest",
   "shardkey": "9",
   "sm_id": 99,
   "date_created": "2021-11-26T06:22:19Z",
   "oof_shard": "1"
}`

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
		kgo.AllowAutoTopicCreation(),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer cl.Close()

	if err := cl.BeginTransaction(); err != nil {
		log.Fatal("can't begin the thansaction: ", err)
	}

	rec := &kgo.Record{
		Topic: topic,
		Key:   []byte("sdfgsdfg"),
		Value: []byte(orderJSON),
		Headers: []kgo.RecordHeader{
			{Key: "content-type", Value: []byte("application/json")},
			{Key: "event-type", Value: []byte("order.created")},
			{Key: "schema-version", Value: []byte("v1")},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if res := cl.ProduceSync(ctx, rec); res.FirstErr() != nil {
		_ = cl.EndTransaction(ctx, false)
		log.Fatal("producer: ", res.FirstErr())
	}

	if err := cl.EndTransaction(context.Background(), true); err != nil {
		log.Fatal("commit tx:", err)
	}

	log.Println("Tx is successful")

}
