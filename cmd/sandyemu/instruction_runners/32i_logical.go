package instruction_runners

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/LamkasDev/sandy/cmd/common/cpu"
	"github.com/LamkasDev/sandy/cmd/common/instruction"
	"github.com/LamkasDev/sandy/cmd/common/util"
	"github.com/LamkasDev/sandy/cmd/sandyemu/internal"
	"github.com/jwalton/gchalk"
	"github.com/rs/zerolog/log"
)

func ProcessORI(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Target] = internal.CPU.Registers.General[ins.Source] | ins.Immediate
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s | %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Target]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Target])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", ins.Immediate)),
		)
	}
}

func ProcessSLL(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Destination] = internal.CPU.Registers.General[ins.Target] << ins.Immediate
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s << %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Destination])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessOR(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Destination] = internal.CPU.Registers.General[ins.Source] | internal.CPU.Registers.General[ins.Target]
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s | %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Destination])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
		)
	}
}

func ProcessSLTU(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	if internal.CPU.Registers.General[ins.Source] < internal.CPU.Registers.General[ins.Target] {
		internal.CPU.Registers.NextGeneral[ins.Destination] = 1
	} else {
		internal.CPU.Registers.NextGeneral[ins.Destination] = 0
	}
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s < %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Destination])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
		)
	}
}

func ProcessANDI(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Target] = internal.CPU.Registers.General[ins.Source] & ins.Immediate
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s & %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Target]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Target])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", ins.Immediate)),
		)
	}
}

func ProcessAND(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Destination] = internal.CPU.Registers.General[ins.Source] & internal.CPU.Registers.General[ins.Target]
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s & %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Destination])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
		)
	}
}

func ProcessSLTI(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	if arch.PlaystationRegister(internal.CPU.Registers.General[ins.Source]) < arch.PlaystationRegister(util.SignExtendRegisterU(ins.Immediate)) {
		internal.CPU.Registers.NextGeneral[ins.Target] = 1
	} else {
		internal.CPU.Registers.NextGeneral[ins.Target] = 0
	}
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s < %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Target]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Target])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", util.SignExtendRegisterU(ins.Immediate))),
		)
	}
}

func ProcessSRA(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Destination] = arch.PlaystationRegisterU(arch.PlaystationRegister(internal.CPU.Registers.General[ins.Target]) >> ins.Immediate)
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s >> %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Destination])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessMFLO(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Load = &cpu.PlaystationCPULoad{
		Register: ins.Destination,
		Value:    internal.CPU.Registers.MultiplyLow,
	}
	if arch.SandyDebug {
		PrintLoad(internal.CPU.Load.Register, internal.CPU.Load.Value)
	}
}

func ProcessSRL(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Destination] = internal.CPU.Registers.General[ins.Target] >> ins.Immediate
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s >> %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Destination])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
			gchalk.Magenta(fmt.Sprint(ins.Immediate)),
		)
	}
}

func ProcessSLTIU(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	if internal.CPU.Registers.General[ins.Source] < util.SignExtendRegisterU(ins.Immediate) {
		internal.CPU.Registers.NextGeneral[ins.Target] = 1
	} else {
		internal.CPU.Registers.NextGeneral[ins.Target] = 0
	}
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s < %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Target]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Target])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", util.SignExtendRegisterU(ins.Immediate))),
		)
	}
}

func ProcessMFHI(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Load = &cpu.PlaystationCPULoad{
		Register: ins.Destination,
		Value:    internal.CPU.Registers.MultiplyHigh,
	}
	if arch.SandyDebug {
		PrintLoad(internal.CPU.Load.Register, internal.CPU.Load.Value)
	}
}

func ProcessSLT(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	if arch.PlaystationRegister(internal.CPU.Registers.General[ins.Source]) < arch.PlaystationRegister(internal.CPU.Registers.General[ins.Target]) {
		internal.CPU.Registers.NextGeneral[ins.Destination] = 1
	} else {
		internal.CPU.Registers.NextGeneral[ins.Destination] = 0
	}
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s < %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Destination])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
		)
	}
}
