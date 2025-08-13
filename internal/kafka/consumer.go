package kafka

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/ummuys/level_0/internal/cache"
	"github.com/ummuys/level_0/internal/service"
	"github.com/ummuys/level_0/internal/validation"
)

func CheckMap(ctx context.Context, logger *zerolog.Logger, cache *cache.OrderCache, orderUID string, info []byte) {
	info, err := cache.Get(ctx, orderUID)
	if err != nil {
		logger.
			Err(err).
			Msg("problem with map")

	}
	logger.Info().
		Str("info", string(info)).
		Msg("all good")

}

func processOrder(pCtx context.Context, orderRawData []byte, orderService service.OrderService) (string, error) {

	order, err := validation.DecodeOrder(orderRawData)

	if err != nil {
		return "", fmt.Errorf("decode err: %w", err)
	}

	UID, err := validation.Validate(order)
	if err != nil {
		return "", fmt.Errorf("validate err: %w", err)
	}

	if err = orderService.Create(pCtx, orderRawData); err != nil {
		return "", err
	}

	return UID, nil
}

func Kafka(pCtx context.Context, logger *zerolog.Logger, orderService service.OrderService, cache *cache.OrderCache) error {

	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")
	group := os.Getenv("KAFKA_GROUP")

	if broker == "" || topic == "" || group == "" {
		logger.Error().
			Str("KAFKA_BROKER", broker).
			Str("KAFKA_TOPIC", topic).
			Str("KAFKA_GROUP", group).
			Msg("missing required env vars")
		return errors.New("kafka: KAFKA_BROKER, KAFKA_TOPIC, KAFKA_GROUP are required")
	}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(broker),
		kgo.ConsumerGroup(group),
		kgo.ConsumeTopics(topic),
		kgo.DisableAutoCommit(),
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

		fetches := cl.PollFetches(pCtx)

		if pCtx.Err() != nil {
			logger.Info().Msg("close kafka routine")
			break
		}

		for _, fe := range fetches.Errors() {
			if !errors.Is(fe.Err, context.Canceled) {
				logger.Error().Err(fe.Err).
					Str("topic", fe.Topic).
					Int32("partition", fe.Partition).
					Msg("fetch error")
			}
		}

		//TODO: сделать красивый коммит
		var hadErr bool
		for _, rec := range fetches.Records() {
			logger.Info().
				Str("key", string(rec.Key)).
				Msg("catch new order")
			if orderKey, err := processOrder(pCtx, rec.Value, orderService); err != nil {
				hadErr = true
				logger.Error().
					Err(err).
					Str("key", string(rec.Key)).
					Msg("can't create a order")

			} else {
				cache.Set(orderKey, rec.Value)
				logger.Info().
					Str("key", string(rec.Key)).
					Msg("order successfuly created")
				CheckMap(pCtx, logger, cache, orderKey, rec.Value)
			}

		}

		if !hadErr {
			if err := cl.CommitUncommittedOffsets(pCtx); err != nil {
				logger.Error().
					Err(err).
					Msg("commit failed")
			}
		}
	}

	fCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := cl.CommitUncommittedOffsets(fCtx); err != nil {
		logger.Warn().Err(err).Msg("final commit failed")
	}
	return nil
}
