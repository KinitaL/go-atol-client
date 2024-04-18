# atol_client

Go клиент для [АТОЛ](https://www.atol.ru)

## Usage

Создание клиента
```go
client := atol_client.NewATOLHttpClient(
    <URL>,
    <LOGIN>,
    <PASSWORD>,
    <GROUP_CODE>,
    <API_VERSION>,
    <RETRY_COUNT>, // кол-во повторных попыток сделать http запрос в случае ошибки 
)
```

Получение чека
```go
client.GetReceipt(&atol_client.GetReceiptRequestMessage{
	UUID: <UUID>,
})
```

Создание чека
```go
resp, err := client.PostReceipt(&atol_client.PostReceiveRequestMessage{
    Operation: "sell", // тип операции по созданию чека (см. документацию ATOL)
    ... // другие параметры в зависимости от типа операции
    },
})
```