#PATHS

PRODUCER_APP_PATH=cmd/producer/main.go
CONSUMER_APP_PATH=cmd/consumer/main.go

DC_DATABASE_PATH=database/docker-compose.yaml
DC_KAFKA_PATH=kafka/docker-compose.yaml



#------------------------------#
.PHONY: run-producer-app run-consumer-app

run-producer-app:
	go run $(PRODUCER_APP_PATH)

run-consumer-app: 
	go run $(CONSUMER_APP_PATH)


#------------------------------#
.PHONY: up-all-containers down-all-containers up-database-container up-kafka-container down-database-container down-kafka-container

down-all-containers: down-database-container down-kafka-container
up-all-containers: up-database-container up-kafka-container

up-database-container:
	docker compose -f $(DC_DATABASE_PATH) up -d 

up-kafka-container:
	docker compose -f $(DC_KAFKA_PATH) up -d 

down-database-container:
	docker compose -f $(DC_DATABASE_PATH) down

down-kafka-container:
	docker compose -f $(DC_KAFKA_PATH) down


	

#------------------------------#