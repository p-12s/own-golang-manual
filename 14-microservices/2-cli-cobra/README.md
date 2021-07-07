# CLI-frameworks (Cobra, urfave/cli)

Примечательно, что kubectl написан с использованием Cobra  
Видео простого примера: https://www.youtube.com/watch?v=7U12a-TTtfo
Репо: https://github.com/spf13/cobra

Запуск текущего примера:
```
// просмотре хелпа
go run main.go help echo
go run main.go help echo times

// использование кастомной команды "echo times"
go run main.go echo times Hello world
go run main.go echo times Hello world -t 3
```
Но на этом все, не думаю что когда-то буду такое использовать - врядли я буду писать тулзу, которая будет принимать из консоли сложные сценарии, как kubectl.  
