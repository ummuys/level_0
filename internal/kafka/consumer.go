package kafka

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/ummuys/level_0/internal/service"
	"github.com/ummuys/level_0/internal/validation"
)

func processOrder(pCtx context.Context, orderRawData []byte, orderService service.OrderService) error {

	order, err := validation.DecodeOrder(orderRawData)
	if err != nil {
		return err
	}

	if err = orderService.Create(pCtx, orderRawData, order); err != nil {
		return err
	}
	return nil
}

func Kafka(pCtx context.Context, logger *zerolog.Logger, orderService service.OrderService) error {

	broker, topic, group, dlqTopic, err := validation.ParseKfkEnv()
	if err != nil {
		return err
	}

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
		kgo.FetchIsolationLevel(kgo.ReadCommitted()),
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

		for _, rec := range fetches.Records() {
			logger.Info().
				Str("key", string(rec.Key)).
				Msg("catch new order")
			if err := processOrder(pCtx, rec.Value, orderService); err != nil {
				logger.Error().
					Err(err).
					Str("key", string(rec.Key)).
					Msg("can't create a order")
				toDLQ(pCtx, logger, cl, dlqTopic, rec, err)
			} else {
				logger.Info().
					Str("key", string(rec.Key)).
					Msg("order successfuly created")
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

func toDLQ(pCtx context.Context, logger *zerolog.Logger, cl *kgo.Client, dlqTopic string, rec *kgo.Record, err error) {

	hdrs := append([]kgo.RecordHeader{}, rec.Headers...) // копия
	hdrs = append(hdrs,
		kgo.RecordHeader{Key: "error", Value: []byte(err.Error())},
		kgo.RecordHeader{Key: "source_topic", Value: []byte(rec.Topic)},
		kgo.RecordHeader{Key: "source_partition", Value: []byte(strconv.Itoa(int(rec.Partition)))},
		kgo.RecordHeader{Key: "source_offset", Value: []byte(strconv.FormatInt(rec.Offset, 10))},
		kgo.RecordHeader{Key: "failed_at", Value: []byte(time.Now().UTC().Format(time.RFC3339))},
	)

	dlqRec := &kgo.Record{
		Topic:     dlqTopic,
		Key:       rec.Key,
		Value:     rec.Value, // оригинальный payload
		Headers:   hdrs,
		Timestamp: time.Now(),
	}

	ctx, cancel := context.WithTimeout(pCtx, 5*time.Second)
	defer cancel()

	if perr := cl.ProduceSync(ctx, dlqRec).FirstErr(); perr != nil {
		logger.Error().Err(perr).
			Str("key", string(rec.Key)).
			Msg("failed to produce to DLQ; will not commit offset (record will be retried)")
		return
	}

	if cerr := cl.CommitRecords(pCtx, rec); cerr != nil {
		logger.Error().
			Err(cerr).
			Msg("failed to commit after DLQ")
	} else {
		logger.Warn().
			Str("key", string(rec.Key)).
			Msg("sent to DLQ and committed")
	}
}
