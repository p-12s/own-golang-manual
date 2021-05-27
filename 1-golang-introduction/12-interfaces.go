package main

import (
	"./gadget"
)

// Player Структуры должны будут иметь как минимум эти методы
type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, song []string) {
	for _, song := range song {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Track")
	player.Stop()

	recorder := player.(gadget.TapeRecorder) // Чтобы вызвать "неинтерфейсный" метод - нужно указать тип с этим кастомным методом
	recorder.Record()
}

func main() {
	var player Player = gadget.TapePlayer{}
	mixtape := []string{
		"First song",
		"Second",
		"Third",
	}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	TryOut(gadget.TapeRecorder{})
}
