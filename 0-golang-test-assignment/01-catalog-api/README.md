## Описание задачи  
  
Необходимо написать простейшее API для каталога товаров. Приложение должно содержать:
- Категории товаров
- Конкретные товары, которые принадлежат к какой-то категории (один товар может принадлежать нескольким категориям)
- Пользователей, которые могут авторизоваться

Возможные действия:
- Получение списка всех категорий
- Получение списка товаров в конкретной категории
- Авторизация пользователей
- Добавление/Редактирование/Удаление категории (для авторизованных пользователей)
- Добавление/Редактирование/Удаление товара (для авторизованных пользователей)

## Технические требования
1. Приложение должно быть написано на Golang
2. Приложение не должно быть написано с помощью какого-либо фреймворка, однако можно устанавливать для него различные пакеты
3. Результаты запросов должны быть представлены в формате JSON
4. Результат задания должен быть выложен на github, должна быть инструкция по запуску проекта. Также необходимо пояснить, сколько на каждую часть проекта ушло времени

## Критерии оценки
- Архитектурная организация API
- Корректная обработка внештатных ситуаций
- Покрытие кода тестами