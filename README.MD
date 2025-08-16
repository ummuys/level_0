# Level 0 for WBTech

Данный проект включается в себя:

    - Consumer: программа, которая считывает сообщения с брокера сообщений, валидирует их и сохраняет в базу + кэш, если они прошли валидацию

    - Producer: программа, которая посылает в брокер сообщение 10 валидных и 10 невалидных заказов

    - Kafka: брокер сообщений. По умолчанию создается два топика. Один для сообщений, другой DLQ. На каждый топик выделено 3 партиции

    - Database: база данных на основе postgresql.

## 🚀 Быстрый старт

### Запуск через make

```bash
cd level_0
make for-inspectors
```

Приложение будет доступно по адресу:
`http://localhost:1337`

Swagger-документация:
`http://localhost:1337/swagger/index.html`

---

## 📂 Структура проекта

```
.
├── Makefile
├── README.MD
├── cmd
│   ├── consumer
│   │   └── main.go
│   └── producer
│       └── main.go
├── docker
│   ├── app-consumer
│   │   └── Dockerfile
│   ├── app-producer
│   │   └── Dockerfile
│   ├── database
│   │   └── migrations
│   │       ├── 20250808190406_create_schema_orders.down.sql
│   │       ├── 20250808190406_create_schema_orders.up.sql
│   │       ├── 20250808190454_create_table_orders_info.down.sql
│   │       ├── 20250808190454_create_table_orders_info.up.sql
│   │       ├── 20250808192910_create_table_payments.down.sql
│   │       ├── 20250808192910_create_table_payments.up.sql
│   │       ├── 20250808192954_create_table_deliveries.down.sql
│   │       ├── 20250808192954_create_table_deliveries.up.sql
│   │       ├── 20250808193039_create_table_items.down.sql
│   │       ├── 20250808193039_create_table_items.up.sql
│   │       ├── 20250812123735_func_insert_order.down.sql
│   │       └── 20250812123735_func_insert_order.up.sql
│   ├── docker-compose.yaml
│   └── kafka
│       └── init-topics.sh
├── go.mod
├── go.sum
├── internal
│   ├── cache
│   │   ├── order_interface.go
│   │   └── order_methods.go
│   ├── config
│   │   ├── cache
│   │   │   ├── cch_config.go
│   │   │   └── cch_env_validate.go
│   │   ├── kafka
│   │   │   ├── kfk_config.go
│   │   │   └── kfk_env_validate.go
│   │   ├── logger
│   │   │   ├── log_config.go
│   │   │   └── log_env_validate.go
│   │   └── server
│   │       ├── srv_config.go
│   │       └── srv_env_validate.go
│   ├── kafka
│   │   └── consumer.go
│   ├── logger
│   │   ├── logfile.go
│   │   └── logger.go
│   ├── models
│   │   ├── orders.go
│   │   ├── orders_db.go
│   │   └── response.go
│   ├── repository
│   │   ├── help.go
│   │   ├── order_interface.go
│   │   ├── order_methods.go
│   │   └── queries.go
│   ├── service
│   │   ├── order_interface.go
│   │   └── order_methods.go
│   ├── testdata
│   │   └── jsons
│   │       ├── bad_json.go
│   │       └── good_json.go
│   ├── validation
│   │   └── validate_order.go
│   └── web
│       ├── endpoints.go
│       ├── handlers
│       │   ├── order_interface.go
│       │   ├── order_methods.go
│       │   ├── server_interface.go
│       │   └── server_methods.go
│       └── server.go
├── logs
│   ├── 2025-08-11.log
│   ├── 2025-08-12.log
│   ├── 2025-08-13.log
│   ├── 2025-08-14.log
│   ├── 2025-08-15.log
│   └── 2025-08-16.log
└── scripts
    ├── check_env.sh
    ├── greetings.sh
    ├── n1.sh
    └── n2.sh
```

---

## 🧩 Используемые технологии

- Go
- net/http
- PostgreSQL
- Swagger (`swaggo`)
- Docker + docker-compose
- zerolog
- Apache Kafka
- RWMutex/Mutex

---

## 📊 Описание эндпоинтов

| Метод | Путь                  | Описание                           | Пример                                     |
| ---------- | ------------------------- | ------------------------------------------ | ------------------------------------------------ |
| POST       | /api/v1/order/{order_uid} | Получить заказ по order_uid | curl http://localhost:1337/api/v1/test_order_uid |

---

## 🐛 Логирование

- Логи хранятся в каталоге `logs/`
- Один файл — на каждый день

---



## 📄 Автор

Ummuys / Егоров Евгений
