package memory

import "github.com/LamkasDev/sandy/cmd/common/arch"

func NewMemoryExpansion2() PlaystationMemorySection {
	section := NewMemorySection(PlaystationMemorySectionTypeExpansion2, 0x1F802000, 0x1F9FFFFF) // 8 kB
	section.GetAddress = func(address arch.PlaystationRegisterU) arch.PlaystationRegisterU {
		if address >= 0xBF802000 {
			return address - 0xBF802000
		}
		if address >= 0x9F802000 {
			return address - 0x9F802000
		}

		return address - section.Start
	}
	section.Matches = func(address arch.PlaystationRegisterU) bool {
		return (address >= section.Start && address <= section.End) || (address >= 0x9F802000 && address <= 0x9F9FFFFF) || (address >= 0xBF802000 && address <= 0xBF9FFFFF)
	}
	section.Write = func(section *PlaystationMemorySection, address arch.PlaystationRegisterU, data []byte) {
		return
	}

	return section
}
