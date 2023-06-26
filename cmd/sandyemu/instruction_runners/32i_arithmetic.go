package instruction_runners

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/LamkasDev/sandy/cmd/common/cpu"
	"github.com/LamkasDev/sandy/cmd/common/instruction"
	"github.com/LamkasDev/sandy/cmd/common/util"
	"github.com/LamkasDev/sandy/cmd/sandyemu/internal"
	"github.com/johncgriffin/overflow"
	"github.com/jwalton/gchalk"
	"github.com/rs/zerolog/log"
)

func ProcessLUI(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Target] = ins.Immediate << 16
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Target]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Target])),
		)
	}
}

func ProcessADDIU(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Target] = internal.CPU.Registers.General[ins.Source] + util.SignExtendRegisterU(ins.Immediate)
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s + %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Target]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Target])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", util.SignExtendRegisterU(ins.Immediate))),
		)
	}
}

func ProcessADDI(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	v, ok := overflow.Add32(int32(internal.CPU.Registers.General[ins.Source]), int32(util.SignExtendRegisterU(ins.Immediate)))
	if !ok {
		log.Error().Msgf("overflow: %v", v)
		panic(1)
	}
	internal.CPU.Registers.NextGeneral[ins.Target] = arch.PlaystationRegisterU(v)
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s + %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Target]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Target])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", util.SignExtendRegisterU(ins.Immediate))),
		)
	}
}

func ProcessADDU(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Destination] = internal.CPU.Registers.General[ins.Source] + internal.CPU.Registers.General[ins.Target]
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s + %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Destination])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
		)
	}
}

func ProcessADD(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	v, ok := overflow.Add32(int32(internal.CPU.Registers.General[ins.Source]), int32(internal.CPU.Registers.General[ins.Target]))
	if !ok {
		log.Error().Msgf("overflow: %v", v)
		panic(1)
	}
	internal.CPU.Registers.NextGeneral[ins.Destination] = arch.PlaystationRegisterU(v)
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s + %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Destination])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
		)
	}
}

func ProcessSUBU(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Destination] = internal.CPU.Registers.General[ins.Source] - internal.CPU.Registers.General[ins.Target]
	if arch.SandyDebug {
		log.Debug().Msgf("%s = %s (%s - %s)",
			gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.NextGeneral[ins.Destination])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
			gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Target])),
		)
	}
}

func PrintDIV(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	log.Debug().Msgf("%s = %s; %s = %s (%s / %s);",
		gchalk.Yellow("mulhi"),
		gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.MultiplyHigh)),
		gchalk.Yellow("mullo"),
		gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.MultiplyLow)),
		gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
		gchalk.Magenta(fmt.Sprintf("%#08x", internal.CPU.Registers.General[ins.Source])),
	)
}

func ProcessDIV(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	n := arch.PlaystationRegister(internal.CPU.Registers.General[ins.Source])
	d := arch.PlaystationRegister(internal.CPU.Registers.General[ins.Target])
	if d == 0 {
		internal.CPU.Registers.MultiplyHigh = arch.PlaystationRegisterU(n)
		if n >= 0 {
			internal.CPU.Registers.MultiplyLow = 0xffffffff
		} else {
			internal.CPU.Registers.MultiplyLow = 1
		}
	} else {
		v, ok := overflow.Div32(int32(n), int32(d))
		if !ok {
			internal.CPU.Registers.MultiplyHigh = 0
			internal.CPU.Registers.MultiplyLow = 0x80000000
		} else {
			internal.CPU.Registers.MultiplyHigh = arch.PlaystationRegisterU(n % d)
			internal.CPU.Registers.MultiplyLow = arch.PlaystationRegisterU(v)
		}
	}

	if arch.SandyDebug {
		PrintDIV(internal, ins)
	}
}

func ProcessDIVU(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	n := internal.CPU.Registers.General[ins.Source]
	d := internal.CPU.Registers.General[ins.Target]
	if d == 0 {
		internal.CPU.Registers.MultiplyHigh = n
		internal.CPU.Registers.MultiplyLow = 0xffffffff
	} else {
		internal.CPU.Registers.MultiplyHigh = n % d
		internal.CPU.Registers.MultiplyLow = n / d
	}

	if arch.SandyDebug {
		PrintDIV(internal, ins)
	}
}
