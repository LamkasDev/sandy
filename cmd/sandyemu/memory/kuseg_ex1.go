package memory

import "github.com/LamkasDev/sandy/cmd/common/arch"

func NewMemoryExpansion1() PlaystationMemorySection {
	section := NewMemorySection(PlaystationMemorySectionTypeExpansion1, 0x1F000000, 0x1F7FFFFF) // 8192 kB
	section.Data = make([]byte, section.Size)
	section.GetAddress = func(address arch.PlaystationRegisterU) arch.PlaystationRegisterU {
		if address >= 0xBF000000 {
			return address - 0xBF000000
		}
		if address >= 0x9F000000 {
			return address - 0x9F000000
		}

		return address - section.Start
	}
	section.Matches = func(address arch.PlaystationRegisterU) bool {
		return (address >= section.Start && address <= section.End) || (address >= 0x9F000000 && address <= 0x9F7FFFFF) || (address >= 0xBF000000 && address <= 0xBF7FFFFF)
	}
	section.Read = func(section *PlaystationMemorySection, address, size arch.PlaystationRegisterU) []byte {
		switch size {
		case 1:
			return []byte{0xff}
		case 2:
			return []byte{0xff, 0xff}
		case 4:
			return []byte{0xff, 0xff, 0xff, 0xff}
		case 8:
			return []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
		}

		return []byte{0xff}
	}

	return section
}
