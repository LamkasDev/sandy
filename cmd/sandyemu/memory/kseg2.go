package memory

import "github.com/LamkasDev/sandy/cmd/common/arch"

func NewKseg2() PlaystationMemorySection {
	section := NewMemorySection(PlaystationMemorySectionTypeKseg2, 0xC0000000, 0xFFFFFFFF) // 1 GB
	section.Write = func(section *PlaystationMemorySection, address arch.PlaystationRegisterU, data []byte) {
		return
	}

	return section
}
