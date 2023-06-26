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

func PrintBranch(counter arch.PlaystationRegisterU) {
	log.Debug().Msgf("%s = %s",
		gchalk.Yellow("c"),
		gchalk.Magenta(fmt.Sprintf("%#08x", counter)),
	)
}

func ProcessJ(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.Counter = (internal.CPU.Registers.Counter & 0xf0000000) | (ins.Immediate << 2)
	if arch.SandyDebug {
		PrintBranch(internal.CPU.Registers.Counter)
	}
}

func ProcessBNE(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	if internal.CPU.Registers.General[ins.Source] != internal.CPU.Registers.General[ins.Target] {
		internal.CPU.Registers.Counter += util.SignExtendRegisterU(ins.Immediate<<2) - 4
	}
	if arch.SandyDebug {
		PrintBranch(internal.CPU.Registers.Counter)
	}
}

func ProcessJAL(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[cpu.PlaystationRegisterReturnAddress] = internal.CPU.Registers.Counter
	ProcessJ(internal, ins)
}

func ProcessJR(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.Counter = internal.CPU.Registers.General[ins.Source]
	if arch.SandyDebug {
		PrintBranch(internal.CPU.Registers.Counter)
	}
}

func ProcessBEQ(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	if internal.CPU.Registers.General[ins.Source] == internal.CPU.Registers.General[ins.Target] {
		internal.CPU.Registers.Counter += util.SignExtendRegisterU(ins.Immediate<<2) - 4
	}
	if arch.SandyDebug {
		PrintBranch(internal.CPU.Registers.Counter)
	}
}

func ProcessBGTZ(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	if internal.CPU.Registers.General[ins.Source] > 0 {
		internal.CPU.Registers.Counter += util.SignExtendRegisterU(ins.Immediate<<2) - 4
	}
	if arch.SandyDebug {
		PrintBranch(internal.CPU.Registers.Counter)
	}
}

func ProcessBLEZ(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	if internal.CPU.Registers.General[ins.Source] <= 0 {
		internal.CPU.Registers.Counter += util.SignExtendRegisterU(ins.Immediate<<2) - 4
	}
	if arch.SandyDebug {
		PrintBranch(internal.CPU.Registers.Counter)
	}
}

func ProcessJALR(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	internal.CPU.Registers.NextGeneral[ins.Destination] = internal.CPU.Registers.Counter
	internal.CPU.Registers.Counter = internal.CPU.Registers.General[ins.Source]
	if arch.SandyDebug {
		PrintBranch(internal.CPU.Registers.Counter)
	}
}

func ProcessBZ(internal *internal.PlaystationInternal, ins instruction.PlaystationInstruction) {
	isBgez := ins.Target&1 > 0
	isLink := ins.Target&16 > 0

	if isLink {
		internal.CPU.Registers.NextGeneral[cpu.PlaystationRegisterReturnAddress] = internal.CPU.Registers.Counter
	}
	if isBgez {
		if internal.CPU.Registers.General[ins.Source] >= 0 {
			internal.CPU.Registers.Counter += util.SignExtendRegisterU(ins.Immediate<<2) - 4
		}
	} else {
		if internal.CPU.Registers.General[ins.Source] < 0 {
			internal.CPU.Registers.Counter += util.SignExtendRegisterU(ins.Immediate<<2) - 4
		}
	}
	if arch.SandyDebug {
		PrintBranch(internal.CPU.Registers.Counter)
	}
}
