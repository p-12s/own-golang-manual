package sample

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD", "Apple M1", "Эльбрус")
}

func randomCPUName() string {
	return randomStringFromSet("i3", "i5", "i7", "i9")
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName() string {
	return randomStringFromSet("RTX", "RX")
}

func randomLaptopName() string {
	return randomStringFromSet("DELL", "HUAWEY", "ASUS", "APPLE")
}

func randomStringFromSet(list ...string) string {
	if len(list) == 0 {
		return ""
	}
	return list[rand.Intn(len(list))]
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomID() string {
	return uuid.New().String()
}
