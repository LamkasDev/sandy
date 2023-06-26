package internal

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/instruction"
	"github.com/jwalton/gchalk"
	"github.com/rs/zerolog/log"
)

func PrintInternalInstruction(ins instruction.PlaystationInstruction) {
	log.Debug().Msgf(
		"c: %s; %s",
		gchalk.Red(fmt.Sprintf("%#06x", ins.Counter)),
		instruction.PrintDecodedInstruction(ins),
	)
}
