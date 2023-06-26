package memory

import "github.com/LamkasDev/sandy/cmd/common/arch"

func NewMemoryBiosRom() PlaystationMemorySection {
	section := NewMemorySection(PlaystationMemorySectionTypeBiosROM, 0x1FC00000, 0x1FCFFFFF) // 512 kB
	section.Data = make([]byte, section.Size)
	section.GetAddress = func(address arch.PlaystationRegisterU) arch.PlaystationRegisterU {
		if address >= 0xBFC00000 {
			return address - 0xBFC00000
		}
		if address >= 0x9FC00000 {
			return address - 0x9FC00000
		}

		return address - section.Start
	}
	section.Matches = func(address arch.PlaystationRegisterU) bool {
		return (address >= section.Start && address <= section.End) || (address >= 0x9FC00000 && address <= 0x9FCFFFFF) || (address >= 0xBFC00000 && address <= 0xBFCFFFFF)
	}

	return section
}
