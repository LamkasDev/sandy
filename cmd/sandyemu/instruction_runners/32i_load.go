package instruction_runners

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/LamkasDev/sandy/cmd/common/cpu"
	"github.com/LamkasDev/sandy/cmd/common/instruction"
	"github.com/LamkasDev/sandy/cmd/common/util"
	"github.com/LamkasDev/sandy/cmd/sandyemu/internal"
	"github.com/LamkasDev/sandy/cmd/sandyemu/memory"
	"github.com/jwalton/gchalk"
	"github.com/rs/zerolog/log"
)

func PrintLoad(target cpu.PlaystationRegisterId, value arch.PlaystationRegisterU) {
	log.Debug().Msgf("%s = %s",
		gchalk.Yellow(cpu.PlaystationRegisterNames[target]),
		gchalk.Magenta(fmt.Sprint(value)),
	)
}

func ProcessLW(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	address := internal.CPU.Registers.General[ins.Source] + util.SignExtendRegisterU(ins.Immediate)
	if internal.Memory.Isolate {
		log.Warn().Msgf("ignoring read while isolate (addr: %04x)", address)
		return
	}
	internal.CPU.Load = &cpu.PlaystationCPULoad{
		Register: ins.Target,
		Value:    arch.PlaystationRegisterU(memory.ReadMemoryWord(&internal.Memory, address)),
	}
	if arch.SandyDebug {
		PrintLoad(internal.CPU.Load.Register, internal.CPU.Load.Value)
	}
}

func ProcessLB(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	address := internal.CPU.Registers.General[ins.Source] + util.SignExtendRegisterU(ins.Immediate)
	if internal.Memory.Isolate {
		log.Warn().Msgf("ignoring read while isolate (addr: %04x)", address)
		return
	}
	internal.CPU.Load = &cpu.PlaystationCPULoad{
		Register: ins.Target,
		Value:    util.SignExtendByte(memory.ReadMemoryWordByte(&internal.Memory, address)),
	}
	if arch.SandyDebug {
		PrintLoad(internal.CPU.Load.Register, internal.CPU.Load.Value)
	}
}

func ProcessLBU(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	address := internal.CPU.Registers.General[ins.Source] + util.SignExtendRegisterU(ins.Immediate)
	if internal.Memory.Isolate {
		log.Warn().Msgf("ignoring read while isolate (addr: %04x)", address)
		return
	}
	internal.CPU.Load = &cpu.PlaystationCPULoad{
		Register: ins.Target,
		Value:    arch.PlaystationRegisterU(arch.PlaystationWordByteU(memory.ReadMemoryWordByte(&internal.Memory, address))),
	}
	if arch.SandyDebug {
		PrintLoad(internal.CPU.Load.Register, internal.CPU.Load.Value)
	}
}
