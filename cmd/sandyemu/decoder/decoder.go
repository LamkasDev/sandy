package decoder

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/LamkasDev/sandy/cmd/common/cpu"
	"github.com/LamkasDev/sandy/cmd/common/instruction"
	"github.com/LamkasDev/sandy/cmd/common/instructions"
	"github.com/LamkasDev/sandy/cmd/common/util"
	"github.com/LamkasDev/sandy/cmd/sandyemu/serrors"
	"github.com/jwalton/gchalk"
)

func Decode(table instructions.PlaystationInstructionTable, raw arch.PlaystationInstruction, counter arch.PlaystationRegisterU) (instruction.PlaystationInstruction, error) {
	code := (uint8)(raw >> 26)
	subtable, ok := table[code]
	if !ok {
		return instruction.PlaystationInstruction{}, fmt.Errorf("%w: %s", serrors.ErrorUnknownInstruction, gchalk.Red(fmt.Sprintf("%#04x", code)))
	}
	format := instruction.PlaystationInstructionFormats[subtable.Format]
	argument := util.SelectRangeRegister(raw, format.Argument.Start, format.Argument.Length)
	entry, ok := subtable.Subtable[(uint8)(argument)]
	if !ok {
		return instruction.PlaystationInstruction{}, fmt.Errorf("%w: %s", serrors.ErrorUnknownInstruction, gchalk.Red(fmt.Sprintf("%#04x", code)))
	}

	return instruction.PlaystationInstruction{
		Definition:  entry.Instruction,
		Process:     entry.Process,
		Counter:     counter,
		Source:      cpu.PlaystationRegisterId(util.SelectRangeRegister(raw, format.Source.Start, format.Source.Length)),
		Target:      cpu.PlaystationRegisterId(util.SelectRangeRegister(raw, format.Target.Start, format.Target.Length)),
		Destination: cpu.PlaystationRegisterId(util.SelectRangeRegister(raw, format.Destination.Start, format.Destination.Length)),
		Immediate:   util.SelectRangeRegister(raw, format.Immediate.Start, format.Immediate.Length),
	}, nil
}
