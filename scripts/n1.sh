#!/bin/sh

clear
echo "Все контейнеры запущены\n"
echo "Теперь перейдем к самому приложению\n"
printf '\033[32m%s\033[0m\n' "Нажми enter, чтобы продолжить"
IFS= read -r _

echo "Сейчас экран очиститься и будет вывод логов из consumer. . ."
sleep 5
clear