package kafka

import (
	"context"
	"errors"
	"os"

	"github.com/rs/zerolog"
	"github.com/twmb/franz-go/pkg/kgo"
)

func Kafka(ctx context.Context, logger *zerolog.Logger) error {

	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")
	group := os.Getenv("KAFKA_GROUP")

	if broker == "" || topic == "" || group == "" {
		logger.Error().
			Str("KAFKA_BROKERS", broker).
			Str("KAFKA_TOPIC", topic).
			Str("KAFKA_GROUP", group).
			Msg("missing required env vars")
		return errors.New("kafka: KAFKA_BROKERS, KAFKA_TOPIC, KAFKA_GROUP are required")
	}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(broker),
		kgo.ConsumerGroup(group),
		kgo.ConsumeTopics(topic),
	)

	if err != nil {
		logger.Error().
			Err(err).
			Msg("can't create a client: ")
		return err
	}

	defer cl.Close()

	logger.Info().
		Str("broker", broker).
		Str("topic", topic).
		Str("group", group).
		Msg("listening")

	for {

		fetches := cl.PollFetches(ctx)

		if ctx.Err() != nil {
			logger.Info().Msg("close kafka routine")
			return nil
		}

		for _, fe := range fetches.Errors() {
			if !errors.Is(fe.Err, context.Canceled) {
				logger.Error().Err(fe.Err).
					Str("topic", fe.Topic).
					Int32("partition", fe.Partition).
					Msg("fetch error")
			}
		}

		for _, rec := range fetches.Records() {
			logger.Info().
				Str("topic", rec.Topic).
				Int32("partition", rec.Partition).
				Int64("offset", rec.Offset).
				Str("key", string(rec.Key)).
				Bytes("value", rec.Value).
				Msg("new msg")
		}
	}
}
