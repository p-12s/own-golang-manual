# Пример сервиса биржевых котировок с websocket и gRPC
[ссылка на видео](https://www.youtube.com/watch?v=OsXjDZ52dyQ)  
[ссылыка на репо](https://github.com/forward32/quotes)  
  
Что будет делать сервис:  
- забирать у биржи L2 order book и хранить актуальное состояние  
- стримить по gRPC срез данных (указывается величина среза и периодичность обновлений)  
![Скрин биржи](https://github.com/p-12s/own-golang-manual/blob/master/8-protobuf-grpc/otus-quotation-exchange/screen.png?raw=true)  

Используем:  
- go 1.16   
- gRPC 3 APIv2 (google.golang.org)   
- websocket  
- evans, wscat для отладки   
  
Выполнение:
![Схема работы](https://github.com/p-12s/own-golang-manual/blob/master/8-protobuf-grpc/otus-quotation-exchange/exchange.png?raw=true)

1. Идем в документацию биржи, смотрим как с ней работать:    
https://docs.onederx.com/#l2-channel  
Пример:  
Сначала отправляем запрос и говорим на какой канал хотим подписаться (в нашем случае l2)  
и на какой инструмент (в нашем случае BTCUSD_P)  
```
{
  "type": "subscribe",
  "payload": {
    "subscriptions": [
      {
        "params": {
          "symbol": "BTCUSD_P"
        },
        "channel": "l2"
      }
    ]
  }
}
```
приходит ответ -   
сначала snapshot (все данные по стакану и, возможно, первый update), потом по-одному update-ы:  
```
{
  "params": {
    "symbol": "BTCUSD_P"
  },
  "type": "snapshot",
  "payload": {
    "snapshot": [
      {
        "count": 1,
        "seq_num": 3996130,
        "timestamp": 1543951133742313500,
        "price": "2943",
        "side": "buy",
        "volume": "121",
        "symbol": "BTCUSD_P"
      },
      ...
    ],
    "updates": [
      {
        "count": 1,
        "seq_num": 3996131,
        "timestamp": 1543957400643094800,
        "price": "3872.5",
        "side": "buy",
        "volume": "47",
        "symbol": "BTCUSD_P"
      },
      ...
    ]
  },
  "channel": "l2"
}
{
  "params": {
    "symbol": "BTCUSD_P"
  },
  "type": "update",
  "payload": {
    "count": 1,
    "seq_num": 3996132,
    "timestamp": 1543957401163098000,
    "price": "3872.5",
    "side": "buy",
    "volume": "45",
    "symbol": "BTCUSD_P"
  },
  "channel": "l2"
}
```
Установим утилиту для проверки работы сторонних вебсокетов и проверим, что биржа отдает данные по-websocket:  
```
npm install -g wscat
wscat -c wss://api.onederx.com/v1/ws
```
Убедимся, что приходят "heartbeat" и подпишемся на канал (отправим запрос по сокету):  
```
{"type": "subscribe","payload": {"subscriptions": [{"params": {"symbol": "BTCUSD_P"},"channel": "l2"}]}}
```
Данные приходят, теперь нужно научиться их считывать и строить "стакан" в памяти.  
2. Реализация кода  
3. Проверка работы  
```
make test // tests

make run // запускает получение данных от биржи

evans -p 50051 -r // можем получить данные от сервиса
    show service
    service Quotes
    call GetL2OrderBook (далее вводим данные и получаем результат)
```
4. Тесты  
2 пути использования тестов:    
- "замокать" биржу, тогда если биржа не работает - тесты работоспособны. Одако про изменение API такие тесты не предупредят и не завалятся  
- стучаться реальным запросом на биржу. Не будут работать без доступности биржи, зато всегда актуальны.   
в примере использован 2 вариант.     
5. Что можно улучшить:    
- трекать ошибку при попытке подписки, если передан неверный инструмент  
- при нерабочей биржи клиенту не уходит уведомление, что данные давно не актуальны  
- оптимизация лока - сейчас он один большой на все инструменты  
- клиенты не кешируют на своей стороне данные, каждый раз пересылаем все заново  
