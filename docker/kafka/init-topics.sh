#!/bin/sh

echo "Creating topics..."
/opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 \
    --create --if-not-exists --topic wb-order-dlq-v1 \
    --partitions 3 --replication-factor 1 \
    --config cleanup.policy=delete \
    --config retention.ms=1209600000 \
    --config min.insync.replicas=1
/opt/kafka/bin/kafka-topics.sh --bootstrap-server kafka:9092 \
    --create --if-not-exists --topic wb-order-created-v1 \
    --partitions 3 --replication-factor 1
echo "Topics ready."