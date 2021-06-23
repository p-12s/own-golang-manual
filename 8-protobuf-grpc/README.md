# gRPC
**Open-source фреймворк от Гугла, который стал частью Cloud Native Computation Foundation (CNCF), наряду с такими продуктами как Docker, Kubernetes.**

Что под капотом фреймворка:  
- Protocol Buffers  
- RPC - remote procedure call 
- HTTP/2

Чем хорош фреймворк gRPC:  
- меньше размер передаваемых данных - они в 16-тиричной кодировке, и там только непосредственно данные
- хорош для взаимодействия микросервисов

Чем плох:
- "не человеко-понятен"
- работает только поверх HTTP/2 (HTTPS)

Типы работы gRPC:  
- Единичный поток               (Unary RPC)
- Стриминг с сервера            (Server streaming RPCs)
- Стриминг с клиента            (Client streaming RPCs)
- Двунаправленный стриминг      (Bidirectional streaming RPCs)
![gRPC](https://github.com/p-12s/own-golang-manual/blob/master/8-protobuf-grpc/gRPC.png?raw=true)

## Best practice
- В файлах *.proto нужно добавлять описание, примеры: [google-pubsub](https://github.com/googleapis/googleapis/blob/master/google/pubsub/v1/pubsub.proto) [google-spanner](https://github.com/googleapis/googleapis/blob/master/google/spanner/v1/spanner.proto)
- Для ручного тестирования работы с помощью Evans, можно подключить reflection

## Субъективное мнение про протоколы
- RPC
    - SOAP
        - 👎🏿 98 год, XML многословен, надо знать WSDL, HTTP/1.1
        - 👍👍 WebSocket

    - JSON-RPC
        - 👎🏿 все еще многословен, HTTP/1.1
        - 👍👍👍 WebSocket, OpenRPC (Swagger)
    - gRPC
        - 👎🏿 нечитаемое тело в бинарном виде
        - 👍👍👍 HTTP/2
- REST
    - RESTfulAPI
        - 👎🏿 многословен
        - 👍👍👍 легко читаемый и простой, кеширование, OpenRPC (Swagger)
    - JSON:API
        - 👎🏿 многословен
        - 👍👍 похож на RESTfulAPI, но строже
    - GraphQL
        - 👎🏿 кеширования нет
        - 👍👍👍 подходит для графовых данных, кеширование

## Применение
- Для взаимодействия подконтрольных **(микро)сервисов** м/у собой - RPC (лучше **gRPC**)
- при разработке **API** для сторонних сервисов или **пользователей** - **JSON-RPC**
- при разработке **клиентского приложения/фронтенда** - **GraphQL, RESTfulAPI**
- остальные протоколы - специфичны
- для gRPC есть возможность автоматом сгенерировать REST-API методы с помощью gRPC gateway.
