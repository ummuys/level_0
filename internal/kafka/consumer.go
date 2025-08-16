package kafka

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/twmb/franz-go/pkg/kgo"
	config "github.com/ummuys/level_0/internal/config/kafka"
	"github.com/ummuys/level_0/internal/service"
	"github.com/ummuys/level_0/internal/validation"
)

func processOrder(pCtx context.Context, orderRawData []byte, orderService service.OrderService) (string, error) {

	order, err := validation.DecodeOrder(orderRawData)
	if err != nil {
		return order.OrderUID, err
	}

	if err = orderService.Create(pCtx, orderRawData, order); err != nil {
		return order.OrderUID, err
	}
	return order.OrderUID, nil
}

func Kafka(pCtx context.Context, logger *zerolog.Logger, orderService service.OrderService) error {

	kfkEnv, err := config.ParseKafkaEnv()
	if err != nil {
		return err
	}

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(kfkEnv.Broker),
		kgo.ConsumerGroup(kfkEnv.Group),
		kgo.ConsumeTopics(kfkEnv.Topic),
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
		Str("broker", kfkEnv.Broker).
		Str("topic", kfkEnv.Topic).
		Str("group", kfkEnv.Group).
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

			logger.Debug().
				Str("evt", "consumer.new_message").
				Msg("")

			if orderUID, err := processOrder(pCtx, rec.Value, orderService); err != nil {

				logger.Debug().
					Str("evt", "order.create.fail").
					Str("order_uid", orderUID).
					Msg("")

				logger.Error().
					Err(err).
					Msg("can't create an order")

				logger.Debug().
					Str("evt", "consumer.start.dlq").
					Msg("")

				toDLQ(pCtx, logger, cl, orderUID, kfkEnv.DlqTopic, rec, err)

			} else {

				logger.Info().
					Str("evt", "order.create.ok").
					Str("orderUID", orderUID).
					Msg("")

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

func toDLQ(pCtx context.Context, logger *zerolog.Logger, cl *kgo.Client, orderUID string, dlqTopic string, rec *kgo.Record, err error) {

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
		Value:     rec.Value, // оригинальный payload
		Headers:   hdrs,
		Timestamp: time.Now(),
	}

	ctx, cancel := context.WithTimeout(pCtx, 5*time.Second)
	defer cancel()

	if perr := cl.ProduceSync(ctx, dlqRec).FirstErr(); perr != nil {
		logger.Error().Err(perr).
			Str("order_uid", orderUID).
			Msg("failed to produce to DLQ; will not commit offset (record will be retried)")
		logger.Debug().
			Str("evt", "consumer.start.dlq.fail").
			Str("order_uid", orderUID).
			Msg("")
		return
	}

	if cerr := cl.CommitRecords(pCtx, rec); cerr != nil {
		logger.Error().
			Err(cerr).
			Str("order_uid", orderUID).
			Msg("failed to commit after DLQ")
		logger.Debug().
			Str("evt", "consumer.start.dlq.fail").
			Str("order_uid", orderUID).
			Msg("")
	} else {
		logger.Warn().
			Str("order_uid", orderUID).
			Msg("sent to DLQ and committed")
		logger.Debug().
			Str("order_uid", orderUID).
			Str("evt", "consumer.start.dlq.ok").
			Msg("")
	}
}
