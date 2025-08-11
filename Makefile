#PATHS

PRODUCER_APP_PATH=cmd/producer/main.go
CONSUMER_APP_PATH=cmd/consumer/main.go

DOCKER_COMPOSE_PATH=docker/docker-compose.yaml



#------------------------------#
.PHONY: run-producer-app run-consumer-app

run-producer-app:
	go run $(PRODUCER_APP_PATH)

run-consumer-app: 
	go run $(CONSUMER_APP_PATH)


#------------------------------#
.PHONY: up-all-containers down-all-containers up-database-container up-kafka-container down-database-container down-kafka-container

down-all-containers: 
	docker compose -f $(DOCKER_COMPOSE_PATH) down

up-all-containers:
	docker compose -f $(DOCKER_COMPOSE_PATH) up -d

up-database-container:
	docker compose -f $(DOCKER_COMPOSE_PATH) up -d database 

up-kafka-container:
	docker compose -f $(DC_KAFKA_PATH) up -d kafka

down-database-container:
	docker compose -f $(DOCKER_COMPOSE_PATH) down database

down-kafka-container:
	docker compose -f $(DC_KAFKA_PATH) down kafka


	

#------------------------------#