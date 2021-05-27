package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Player playing ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Player stop ")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Recorder playing ", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Recorder stop ")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recorder record ")
}
