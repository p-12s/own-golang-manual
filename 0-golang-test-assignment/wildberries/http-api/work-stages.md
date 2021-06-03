# Этапы выполения

## Создание основы HTTP-сервиса

## Добавление хендлеров и основных сущностей
- InitRoutes() добавление API-адресов, как они должны быть в задании
  pkg/handler/handler.go
- Создание сигнатур хендлеров
  /pkg/handler/user.go
  /pkg/handler/comment.go
- Добавлены основные сущности
  /user.go
  /http_api.go

## Порядок инициализации: Repository->Service->Handler (DI)








