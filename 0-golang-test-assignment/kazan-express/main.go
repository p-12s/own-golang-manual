package main

import (
	"context"
	"fmt"
	"github.com/own-golang-manual/0-golang-test-assignment/kazan-express/service"
	"log"
	"math"
	"time"
)

func main() {
	// заглушка-ответ от внешнего сервиса
	externalService := service.ExternalService{
		MaxProcessItem:        10,
		ProcessPeriod:         3, // пусть период обработки будет в секундах
		AlreadyProcessedCount: 0, // для простоты примем, что уже обработано столько (сервис брал бы это кол-во из своей бд, сколько он уже обработал за последний период времени)
	}
	// имитируем запрос лимита и периода у внешнего сервиса
	limit, duration := externalService.GetLimits()

	var ctx context.Context // пустой контекст для передачи в метод (дана такая сигнатура)

	// отправляем пачку элементов в кол-ве = n/(t*60 сек) в период времени - для равномерности
	itemInSec := uint64(math.Floor(float64(limit) / float64(duration)))
	batch := service.CreateBatch(itemInSec)

	for range time.Tick(time.Second) {
		err := externalService.Process(ctx, batch)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println("Batch send period 1 sec")
	}
}
