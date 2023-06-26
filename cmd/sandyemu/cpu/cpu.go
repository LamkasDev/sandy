package cpu

import (
	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/LamkasDev/sandy/cmd/common/cpu"
	"github.com/LamkasDev/sandy/cmd/common/instruction"
)

type PlaystationCPURegisters struct {
	General      [32]arch.PlaystationRegisterU
	NextGeneral  [32]arch.PlaystationRegisterU
	Counter      arch.PlaystationRegisterU
	MultiplyHigh arch.PlaystationRegisterU
	MultiplyLow  arch.PlaystationRegisterU
}

type PlaystationCPU struct {
	Registers    PlaystationCPURegisters
	Cache        *instruction.PlaystationInstruction
	CacheCounter arch.PlaystationRegisterU
	Load         *cpu.PlaystationCPULoad
	TotalCycles  uint64
}

func NewPlaystationCPU() PlaystationCPU {
	return PlaystationCPU{
		Registers: PlaystationCPURegisters{
			Counter: arch.PlaystationBootRomVector,
		},
		CacheCounter: arch.PlaystationBootRomVector,
	}
}
