package instruction

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/LamkasDev/sandy/cmd/common/cpu"
	"github.com/jwalton/gchalk"
)

type PlaystationInstruction struct {
	Definition  PlaystationInstructionDefinition
	Process     interface{}
	Counter     arch.PlaystationRegisterU
	Source      cpu.PlaystationRegisterId
	Target      cpu.PlaystationRegisterId
	Destination cpu.PlaystationRegisterId
	Immediate   arch.PlaystationRegisterU
}

func PrintDecodedInstruction(ins PlaystationInstruction) string {
	return fmt.Sprintf(
		"code: %-36s; src: %-15s; trg: %-15s; dst: %-15s; im: %s; ",
		fmt.Sprintf("%s (%s)", gchalk.Green(fmt.Sprintf("%#04x", ins.Definition.Code)), gchalk.Blue(ins.Definition.Name)),
		gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Source]),
		gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Target]),
		gchalk.Yellow(cpu.PlaystationRegisterNames[ins.Destination]),
		gchalk.Magenta(fmt.Sprintf("%#08x", ins.Immediate)),
	)
}
