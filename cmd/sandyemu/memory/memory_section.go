package memory

import (
	"github.com/LamkasDev/sandy/cmd/common/arch"
)

const (
	PlaystationMemorySectionTypeNull       = PlaystationMemorySectionType("NULL")
	PlaystationMemorySectionTypeKuseg      = PlaystationMemorySectionType("KUSEG")
	PlaystationMemorySectionTypeKseg2      = PlaystationMemorySectionType("KSEG2")
	PlaystationMemorySectionTypeBiosROM    = PlaystationMemorySectionType("BIOS_ROM")
	PlaystationMemorySectionTypeIRQControl = PlaystationMemorySectionType("IRQ_CONTROL")
	PlaystationMemorySectionTypeExpansion1 = PlaystationMemorySectionType("EXPANSION_1")
	PlaystationMemorySectionTypeIO         = PlaystationMemorySectionType("IO")
	PlaystationMemorySectionTypeExpansion2 = PlaystationMemorySectionType("EXPANSION_2")
)

type PlaystationMemorySectionType string
type PlaystationMemorySectionMatches func(address arch.PlaystationRegisterU) bool
type PlaystationMemorySectionGetAddress = func(address arch.PlaystationRegisterU) arch.PlaystationRegisterU
type PlaystationMemorySectionRead func(section *PlaystationMemorySection, address arch.PlaystationRegisterU, size arch.PlaystationRegisterU) []byte
type PlaystationMemorySectionWrite func(section *PlaystationMemorySection, address arch.PlaystationRegisterU, data []byte)

type PlaystationMemorySection struct {
	Type PlaystationMemorySectionType

	Start arch.PlaystationRegisterU
	End   arch.PlaystationRegisterU
	Size  arch.PlaystationRegisterU
	Data  []byte

	GetAddress PlaystationMemorySectionGetAddress
	Matches    PlaystationMemorySectionMatches

	Read  PlaystationMemorySectionRead
	Write PlaystationMemorySectionWrite
}

func NewMemorySection(sectionType PlaystationMemorySectionType, start arch.PlaystationRegisterU, end arch.PlaystationRegisterU) PlaystationMemorySection {
	return PlaystationMemorySection{
		Type:       sectionType,
		Start:      start,
		End:        end,
		Size:       (end + 1) - start,
		GetAddress: func(address arch.PlaystationRegisterU) arch.PlaystationRegisterU { return address - start },
		Matches: func(address arch.PlaystationRegisterU) bool {
			return address >= start && address <= end
		},
		Read: func(section *PlaystationMemorySection, address arch.PlaystationRegisterU, size arch.PlaystationRegisterU) []byte {
			return section.Data[address : address+size]
		},
		Write: func(section *PlaystationMemorySection, address arch.PlaystationRegisterU, data []byte) {
			copy(section.Data[address:], data)
		},
	}
}
