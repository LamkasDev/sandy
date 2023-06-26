//go:build sandy32

package memory

import (
	"encoding/binary"

	"github.com/LamkasDev/sandy/cmd/common/arch"
)

func ReadMemoryWordDouble(memory *PlaystationMemory, address arch.PlaystationRegisterU) arch.PlaystationWordDouble {
	return arch.PlaystationWordDouble(binary.LittleEndian.Uint64(ReadMemory(memory, address, 8)))
}

func ReadMemoryWord(memory *PlaystationMemory, address arch.PlaystationRegisterU) arch.PlaystationWord {
	return arch.PlaystationWord(binary.LittleEndian.Uint32(ReadMemory(memory, address, 4)))
}

func ReadMemoryWordHalf(memory *PlaystationMemory, address arch.PlaystationRegisterU) arch.PlaystationWordHalf {
	return arch.PlaystationWordHalf(binary.LittleEndian.Uint16(ReadMemory(memory, address, 2)))
}

func ReadMemoryWordByte(memory *PlaystationMemory, address arch.PlaystationRegisterU) arch.PlaystationWordByte {
	return arch.PlaystationWordByte(ReadMemory(memory, address, 1)[0])
}

func WriteMemoryWordDouble(memory *PlaystationMemory, address arch.PlaystationRegisterU, data arch.PlaystationWordDouble) {
	WriteMemory(memory, address, binary.LittleEndian.AppendUint64([]byte{}, uint64(data)))
}

func WriteMemoryWord(memory *PlaystationMemory, address arch.PlaystationRegisterU, data arch.PlaystationWord) {
	WriteMemory(memory, address, binary.LittleEndian.AppendUint32([]byte{}, uint32(data)))
}

func WriteMemoryWordHalf(memory *PlaystationMemory, address arch.PlaystationRegisterU, data arch.PlaystationWordHalf) {
	WriteMemory(memory, address, binary.LittleEndian.AppendUint16([]byte{}, uint16(data)))
}

func WriteMemoryWordByte(memory *PlaystationMemory, address arch.PlaystationRegisterU, data arch.PlaystationWordByte) {
	WriteMemory(memory, address, []byte{byte(data)})
}
