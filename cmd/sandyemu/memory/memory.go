package memory

import (
	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/rs/zerolog/log"
)

type PlaystationMemory struct {
	Points   map[arch.PlaystationRegisterU]PlaystationMemorySection
	Sections []PlaystationMemorySection
	Isolate  bool
}

func NewMemory() PlaystationMemory {
	return PlaystationMemory{
		Sections: []PlaystationMemorySection{NewMemoryIrqControl(), NewMemoryExpansion1(), NewMemoryIO(), NewMemoryExpansion2(), NewMemoryBiosRom(), NewKuseg(), NewKseg2()},
	}
}

func ReadMemory(memory *PlaystationMemory, address arch.PlaystationRegisterU, size arch.PlaystationRegisterU) []byte {
	point, ok := memory.Points[address]
	if ok {
		return point.Read(&point, address, size)
	}
	for _, section := range memory.Sections {
		if section.Matches(address) {
			return section.Read(&section, section.GetAddress(address), size)
		}
	}

	log.Warn().Msgf("illegal read (addr: %04x)", address)
	return []byte{}
}

func WriteMemory(memory *PlaystationMemory, address arch.PlaystationRegisterU, data []byte) {
	point, ok := memory.Points[address]
	if ok {
		point.Write(&point, address, data)
		return
	}
	for _, section := range memory.Sections {
		if section.Matches(address) {
			section.Write(&section, section.GetAddress(address), data)
			return
		}
	}
	log.Warn().Msgf("illegal write (addr: %04x)", address)
}
