#PATHS

PRODUCER_APP_PATH=cmd/producer/main.go
CONSUMER_APP_PATH=cmd/consumer/main.go

DOCKER_COMPOSE_PATH=docker/docker-compose.yaml

GR=scripts/greetings.sh
N1=scripts/n1.sh
N2=scripts/n2.sh
N3=scripts/n3.sh
N4=scripts/n4.sh

.PHONY: for-inspectors

for-inspectors:
	@chmod +x $(GR) $(N1) $(N2) $(N3) $(N4)

	@./$(GR)
	docker compose -f $(DOCKER_COMPOSE_PATH) up -d

	@./$(N1)
	docker logs app

	@./$(N2)

# 	FIX THIS

	@./$(N3)
	@docker logs --tail 4 app 

	@./$(N4)
	@docker logs --tail 7 app 




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
	docker compose -f $(DOCKER_COMPOSE_PATH) down kafka

#------------------------------#
#FOR DEV. DON'T USE 
docker-restart-all:
	docker compose -f $(DOCKER_COMPOSE_PATH) down -v
	docker image rm docker-app docker-app-producer
#------------------------------#