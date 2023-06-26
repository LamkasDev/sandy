package memory

import (
	"github.com/LamkasDev/sandy/cmd/common/arch"
)

func NewKuseg() PlaystationMemorySection {
	section := NewMemorySection(PlaystationMemorySectionTypeKuseg, 0x00, 0x7FFFFFFF) // 2 GB
	section.Data = make([]byte, section.Size)
	section.GetAddress = func(address arch.PlaystationRegisterU) arch.PlaystationRegisterU {
		if address >= 0xA0000000 {
			return address - 0xA0000000
		}
		if address >= 0x80000000 {
			return address - 0x80000000
		}

		return address - section.Start
	}
	section.Matches = func(address arch.PlaystationRegisterU) bool {
		return address <= 0xBFFFFFFF
	}

	return section
}
