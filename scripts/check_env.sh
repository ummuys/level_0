#!/bin/sh

env_consumer_path=.env.consumer
env_producer_path=.env.producer
env_docker_path=docker/.env

if [ -f "$env_consumer_path" ]; then
  printf '\033[32m%s\033[0m\n' "env файл для consumer найден"
else
  printf '\033[1;31m%s\033[0m\n' "env файл для consumer не найден"
  printf "копирую конфиг с .env.consumer.example\n"
  cp ".env.consumer.example" "$env_consumer_path"
fi

if [ -f "$env_producer_path" ]; then
  printf '\033[32m%s\033[0m\n' "env файл для producer найден"
else
  printf '\033[1;31m%s\033[0m\n' "env файл для producer не найден"
  printf "копирую конфиг с .env.producer.example\n"
  cp ".env.producer.example" "$env_producer_path"
fi

if [ -f "$env_docker_path" ]; then
  printf '\033[32m%s\033[0m\n' "env файл для docker найден"
else
  printf '\033[1;31m%s\033[0m\n' "env файл для docker не найден"
  printf "копирую конфиг с .env.example\n"
  cp "docker/.env.example" "$env_docker_path"
fi
echo "\n"
printf '\033[32m%s\033[0m\n' "Нажми enter, чтобы продолжить"
IFS= read -r _
