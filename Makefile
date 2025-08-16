#PATHS

PRODUCER_APP_PATH=cmd/producer/main.go
CONSUMER_APP_PATH=cmd/consumer/main.go

DOCKER_COMPOSE_PATH=docker/docker-compose.yaml

SCRIPTS_PATH=scripts

GR=$(SCRIPTS_PATH)/greetings.sh
CE=$(SCRIPTS_PATH)/check_env.sh
N1=$(SCRIPTS_PATH)/n1.sh
N2=$(SCRIPTS_PATH)/n2.sh
N3=$(SCRIPTS_PATH)/n3.sh
N4=$(SCRIPTS_PATH)/n4.sh

.PHONY: for-inspectors

for-inspectors:
	@chmod +x $(CE) $(GR) $(N1) $(N2) $(N3) $(N4)

	@./$(CE)

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