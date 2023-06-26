package instruction_runners

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/LamkasDev/sandy/cmd/common/instruction"
	"github.com/LamkasDev/sandy/cmd/common/util"
	"github.com/LamkasDev/sandy/cmd/sandyemu/internal"
	"github.com/LamkasDev/sandy/cmd/sandyemu/memory"
	"github.com/jwalton/gchalk"
	"github.com/rs/zerolog/log"
)

func PrintSave(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	log.Debug().Msgf("mem[%s] = %s",
		gchalk.BrightGreen(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source]+util.SignExtendRegisterU(ins.Immediate))),
		gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
	)
}

func ProcessSW(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	address := internal.CPU.Registers.General[ins.Source] + util.SignExtendRegisterU(ins.Immediate)
	if internal.Memory.Isolate {
		log.Warn().Msgf("ignoring write while isolate (addr: %04x)", address)
		return
	}
	memory.WriteMemoryWord(&internal.Memory, address, arch.PlaystationWord(internal.CPU.Registers.General[ins.Target]))
	if arch.SandyDebug {
		PrintSave(internal, ins)
	}
}

func ProcessSH(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	address := internal.CPU.Registers.General[ins.Source] + util.SignExtendRegisterU(ins.Immediate)
	if internal.Memory.Isolate {
		log.Warn().Msgf("ignoring write while isolate (addr: %04x)", address)
		return
	}
	memory.WriteMemoryWordHalf(&internal.Memory, address, arch.PlaystationWordHalf(internal.CPU.Registers.General[ins.Target]))
	if arch.SandyDebug {
		PrintSave(internal, ins)
	}
}

func ProcessSB(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	address := internal.CPU.Registers.General[ins.Source] + util.SignExtendRegisterU(ins.Immediate)
	if internal.Memory.Isolate {
		log.Warn().Msgf("ignoring write while isolate (addr: %04x)", address)
		return
	}
	memory.WriteMemoryWordByte(&internal.Memory, address, arch.PlaystationWordByte(internal.CPU.Registers.General[ins.Target]))
	if arch.SandyDebug {
		PrintSave(internal, ins)
	}
}
