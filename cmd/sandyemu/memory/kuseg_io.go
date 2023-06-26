package memory

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/jwalton/gchalk"
	"github.com/rs/zerolog/log"
)

func NewMemoryIO() PlaystationMemorySection {
	section := NewMemorySection(PlaystationMemorySectionTypeIO, 0x1F801000, 0x1F801FFF) // 4 kB
	section.GetAddress = func(address arch.PlaystationRegisterU) arch.PlaystationRegisterU {
		if address >= 0xBF801000 {
			return address - 0xBF801000
		}
		if address >= 0x9F801000 {
			return address - 0x9F801000
		}

		return address - section.Start
	}
	section.Matches = func(address arch.PlaystationRegisterU) bool {
		return (address >= section.Start && address <= section.End) || (address >= 0x9F801000 && address <= 0x9F801FFF) || (address >= 0xBF801000 && address <= 0xBF801FFF)
	}
	section.Write = func(section *PlaystationMemorySection, address arch.PlaystationRegisterU, data []byte) {
		return
	}
	section.Read = func(section *PlaystationMemorySection, address, size arch.PlaystationRegisterU) []byte {
		log.Info().Msgf("io read (addr: %s)", gchalk.Red(fmt.Sprint(address)))
		return section.Data[address : address+size]
	}

	return section
}
