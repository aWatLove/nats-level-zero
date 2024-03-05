# WB Tech: level #0
## Task
См. файл: `task.pdf`
## Instruction to getting start
1. Скачайте репозиторий\
`git clone https://github.com/aWatLove/nats-level-zero.git`
2. Смените директорию\
`cd nats-level-zero`
3. Build and run docker `docker compose up --build` или используйте утилиту Make `make build`
4. После запуска всех контейнеров и автомиграции БД, можно протестировать API\
`localhost:8080/swagger/index.html#/`

Чтобы добавить новый Order, можно изменить файл `./static/model.json`, после чего запустить специальный скрипт на Go: `go run github.com/aWatLove/nats-lvl-zero/cmd/publisher`

### Technologies:
1. Golang
2. Gin
3. GORM
4. PostgreSQL
5. Nats-Streaming
6. Docker
7. Swagger

## Примеры запросов
### GET order by uid
`http://localhost:8080/api/order/:uid`

пример uid для запроса: `b563feb7b2b84b6test`

Output:
```json
{
    "order_uid": "b563feb7b2b84b6test",
    "track_number": "WBILMTESTTRACK",
    "entry": "WBIL",
    "delivery": {
        "name": "Test Testov",
        "phone": "+9720000000",
        "zip": "2639809",
        "city": "Kiryat Mozkin",
        "address": "Ploshad Mira 15",
        "region": "Kraiot",
        "email": "test@gmail.com"
    },
    "payment": {
        "transaction": "b563feb7b2b84b6test",
        "request_id": "",
        "currency": "USD",
        "provider": "wbpay",
        "amount": 1817,
        "payment_dt": 1637907727,
        "bank": "alpha",
        "delivery_cost": 1500,
        "goods_total": 317,
        "custom_fee": 0
    },
    "items": [
        {
            "chrt_id": 9934930,
            "track_number": "WBILMTESTTRACK",
            "price": 453,
            "rid": "ab4219087a764ae0btest",
            "name": "Mascaras",
            "sale": 30,
            "size": "0",
            "total_price": 317,
            "nm_id": 2389212,
            "brand": "Vivienne Sabo",
            "status": 202
        }
    ],
    "locale": "en",
    "internal_signature": "",
    "customer_id": "test",
    "delivery_service": "meest",
    "shardkey": "9",
    "sm_id": 99,
    "date_created": "2021-11-26T06:22:19Z",
    "oof_shard": "1"
}
```
## GET order by uid from DB
`http://localhost:8080/api/order/db/:uid`

пример uid для запроса: `b563feb7b2b84b6test`

Output: такой же как и в методе `GET order by uid`

### GET all orders
`http://localhost:8080/api/order`

Output: 
```json
{
    "data": [
        {
            "order_uid": "b563feb7b2b84b6test",
            "track_number": "WBILMTESTTRACK",
            "entry": "WBIL",
            "delivery": {
                "name": "Test Testov",
                "phone": "+9720000000",
                "zip": "2639809",
                "city": "Kiryat Mozkin",
                "address": "Ploshad Mira 15",
                "region": "Kraiot",
                "email": "test@gmail.com"
            },
            "payment": {
                "transaction": "b563feb7b2b84b6test",
                "request_id": "",
                "currency": "USD",
                "provider": "wbpay",
                "amount": 1817,
                "payment_dt": 1637907727,
                "bank": "alpha",
                "delivery_cost": 1500,
                "goods_total": 317,
                "custom_fee": 0
            },
            "items": [
                {
                    "chrt_id": 9934930,
                    "track_number": "WBILMTESTTRACK",
                    "price": 453,
                    "rid": "ab4219087a764ae0btest",
                    "name": "Mascaras",
                    "sale": 30,
                    "size": "0",
                    "total_price": 317,
                    "nm_id": 2389212,
                    "brand": "Vivienne Sabo",
                    "status": 202
                }
            ],
            "locale": "en",
            "internal_signature": "",
            "customer_id": "test",
            "delivery_service": "meest",
            "shardkey": "9",
            "sm_id": 99,
            "date_created": "2021-11-26T06:22:19Z",
            "oof_shard": "1"
        }
    ]
}
```

### Error
В случаях когда запись не найдена, возвращается ошибка

Пример: 500 status code
```json
{
  "message": "record not found"
}
```


## Stress test
### WRK
Тестирование метода `/api/order/db/:uid`
```
  40 goroutine(s) running concurrently
2604 requests in 5.0362152s, 2.35MB read
Requests/sec:           517.05
Transfer/sec:           477.16KB
Avg Req Time:           77.361216ms
Fastest Request:        3.0287ms
Slowest Request:        263.957ms
Number of Errors:       0
```

Тестирование метода `/api/order/:uid`
```
  40 goroutine(s) running concurrently
43186 requests in 4.98855358s, 38.92MB read
Requests/sec:           8657.02
Transfer/sec:           7.80MB
Avg Req Time:           4.620528ms
Fastest Request:        499.8µs
Slowest Request:        736.6561ms
Number of Errors:       0
```
Тестирование метода `/api/order`
```
  40 goroutine(s) running concurrently
43715 requests in 4.98395387s, 74.75MB read
Requests/sec:           8771.15
Transfer/sec:           15.00MB
Avg Req Time:           4.560406ms
Fastest Request:        662.5µs
Slowest Request:        18.0338ms
Number of Errors:       0
```


## Как улучшить сервис?
1. Переосмыслить кеш в сервисе
2. Исключить gorm.Automigrate(), а использовать скрипты миграции для создания таблиц БД