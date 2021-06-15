package sample

import (
	"github.com/p-12s/own-golang-manual/8-protobuf-grpc/tech-school/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewCPU() *pb.CPU {
	numberCores := randomInt(2, 8)
	minGhz := randomFloat64(2.0, 3.5)

	return &pb.CPU{
		Brand:         randomCPUBrand(),
		Name:          randomCPUName(),
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(randomInt(numberCores, 12)),
		MinGhz:        minGhz,
		MaxGhz:        randomFloat64(minGhz, 5.0),
	}
}

func NewGPU() *pb.GPU {
	minGhz := randomFloat64(1.0, 1.5)

	return &pb.GPU{
		Brand:  randomGPUBrand(),
		Name:   randomGPUName(),
		MinGhz: minGhz,
		MaxGhz: randomFloat64(minGhz, 2.0),
		Memory: &pb.Memory{
			Value: uint64(randomInt(2, 6)),
			Uint:  pb.Memory_GIGABYTE,
		},
	}
}

func NewRAM() *pb.Memory {
	return &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Uint:  pb.Memory_GIGABYTE,
	}
}

func NewSSD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Uint:  pb.Memory_GIGABYTE,
		},
	}
}

func NewHDD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Uint:  pb.Memory_TERABYTE,
		},
	}
}

func NewLaptop() *pb.Laptop {
	return &pb.Laptop{
		Id:      randomID(),
		Brand:   randomLaptopName(),
		Cpu:     NewCPU(),
		Gpu:     []*pb.GPU{NewGPU()},
		Ram:     NewRAM(),
		Storage: []*pb.Storage{NewHDD(), NewSSD()},
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2021)),
		UpdatedAt:   timestamppb.Now(),
	}
}
