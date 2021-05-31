package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"time"
	"github.com/own-golang-manual/0-golang-test-assignment/kazan-express/service"
)

func main() {
	externalService := service.ExternalService {
		MaxProcessItem: 100,
		ProcessPeriod: 3, // пусть период обработки будет в секундах
		alreadyProcessedCount: 0, // для простоты примем, что уже обработано столько
	}

	// запрашиваем лимит и период
	limit, duration := externalService.GetLimits()

	var ctx context.Context
	batch1 := createBatch(limit)
	fmt.Println(batch1)
	// 1) отправляем пачку = n с периодичностью t
	for _ = range time.Tick(time.Second * duration) {
		err := externalService.Process(ctx, batch1)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println("Batch type 1 send")
	}
	// минусы подхода - если за предыдущий период (до запуска) счетчик обработанных задач уже > 0 - будет ошибка
	// нужно сначала запросить: "у меня счетчик обработанных Item обнулен? Если нет - уменьшать первый Batch на это количество"

	// 2) отправляем пачку = n/(t*60 сек) в период времени - для равномерности
	itemInSec := uint64(math.Floor(float64(limit)/float64(duration)))
	batch2 := createBatch(itemInSec)

	for _ = range time.Tick(time.Second) {
		err := externalService.Process(ctx, batch2)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println("Batch type 2 send")
	}
}