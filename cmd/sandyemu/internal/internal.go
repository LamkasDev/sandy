package internal

import (
	"github.com/LamkasDev/sandy/cmd/sandyemu/cpu"
	"github.com/LamkasDev/sandy/cmd/sandyemu/gpu"
	"github.com/LamkasDev/sandy/cmd/sandyemu/memory"
)

type PlaystationInternal struct {
	Running   bool
	Executing bool
	CPU       cpu.PlaystationCPU
	CP0       cpu.PlaystationCP0
	Memory    memory.PlaystationMemory
	GPU       gpu.PlaystationGPU
}

func NewInternal() PlaystationInternal {
	return PlaystationInternal{
		Running:   true,
		Executing: true,
		CPU:       cpu.NewPlaystationCPU(),
		CP0:       cpu.NewPlaystationCP0(),
		Memory:    memory.NewMemory(),
		GPU:       gpu.NewPlaystationGPU(),
	}
}
