package instruction_runners

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/LamkasDev/sandy/cmd/common/cpu"
	"github.com/LamkasDev/sandy/cmd/common/instruction"
	"github.com/LamkasDev/sandy/cmd/sandyemu/internal"
	"github.com/jwalton/gchalk"
	"github.com/rs/zerolog/log"
)

func ProcessMTC0(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CP0.Registers[ins.Destination] = internal.CPU.Registers.General[ins.Target]
	switch ins.Destination {
	case cpu.PlaystationCp0RegisterStatus:
		if internal.CP0.Registers[ins.Destination]&cpu.PlaystationCP0StatusIsolateCache != 0 {
			internal.Memory.Isolate = true
		} else {
			internal.Memory.Isolate = false
		}
	}
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s",
			gchalk.Yellow(cpu.PlaystationCp0RegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CP0.Registers[ins.Destination])),
		)
	}
}

func ProcessMFC0(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Load = &cpu.PlaystationCPULoad{
		Register: ins.Target,
		Value:    internal.CP0.Registers[ins.Destination],
	}
	if arch.SandyDebug {
		PrintLoad(internal.CPU.Load.Register, internal.CPU.Load.Value)
	}
}
