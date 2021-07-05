# Инструменты для реализации API-сервера на Golang с автогенерацией кода и документации
[презентация](https://drive.google.com/file/d/1UYrwKaJ0UqnnVzISZ3wVfEl4tD7zJzEG/view)  
[видео](https://www.youtube.com/watch?v=Q9x1FVPDGu4&list=PL4jag8ijtDPyYMb9_JDZ2UYZko2sGBcOj&index=10)    

Swagger ≠ OpenAPI  
[OpenAPI](https://swagger.io/specification/) - спецификация, способ описания АПИ.   
[Swagger](https://editor.swagger.io) - инструмент для генерации/просмотра/редактирования АПИ.

## Инструменты
[swagger-api/swagger-codegen](https://github.com/swagger-api/swagger-codegen)
- 👍 есть веб-версия, в которой можно редактировать АПИ
- 👍 можно генерировать клиент/сервер для разных языков
- 👎🏿 по-большей части это заглушка веб-сервера 
---
[go-swagger/go-swagger](https://github.com/go-swagger/go-swagger)
- 👍 генерация клиента/сервера
- 👍 отправка запросов 
- 👍 возможность генерации документации по специально-подготовленным сигнатурам методов
- 👎🏿 нет поддержки OpenAPI 3.x
- 👎🏿 поддерживает не все функции JSON Schema
- 👎🏿 не валидирует возвращаемые ответы сервера
- 👎🏿 не вадидирует параметры на клиенте  
---
[swaggo/swag](https://github.com/swaggo/swag)
- 👍 поддерживает популярные фреймворки (gin/echo/net-http/fiber/...)
- 👎🏿 нет поддержки OpenAPI 3.x
---
[grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
- 👍 с помощью protobuf генерирует и код и доку, на выходе получаем (RESTFull JSON API/OpenAPI/WebSocket/gRPC)
- 👍 protobuf компактнее JSON/XML/Thrift, менее требователен к ресурсам
- 👎🏿 нет поддержки OpenAPI 3.x
- 👎🏿 не читается по сети (например, в трейсе)
- 👎🏿 нет значений по-умолчанию
---