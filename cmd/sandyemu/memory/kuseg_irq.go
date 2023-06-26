package memory

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/jwalton/gchalk"
	"github.com/rs/zerolog/log"
)

func NewMemoryIrqControl() PlaystationMemorySection {
	section := NewMemorySection(PlaystationMemorySectionTypeIRQControl, 0x1f801070, 0x1f801078)
	section.Data = make([]byte, section.Size)
	section.Write = func(section *PlaystationMemorySection, address arch.PlaystationRegisterU, data []byte) {
		return
	}
	section.Read = func(section *PlaystationMemorySection, address, size arch.PlaystationRegisterU) []byte {
		log.Info().Msgf("irq read (addr: %s)", gchalk.Red(fmt.Sprint(address)))
		return section.Data[address : address+size]
	}

	return section
}
