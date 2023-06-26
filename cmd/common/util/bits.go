package util

import (
	"math"

	"github.com/LamkasDev/sandy/cmd/common/arch"
)

func SelectRangeRegister(raw arch.PlaystationInstruction, start uint8, length uint8) arch.PlaystationRegisterU {
	return arch.PlaystationRegisterU((raw >> start) & (1<<length - 1))
}

func InsertRangeRegister(accumulator arch.PlaystationRegisterU, current arch.PlaystationRegisterU, start uint8, length uint8) arch.PlaystationRegisterU {
	// Zero out the range where we're planning to insert current range
	mask := ^(^arch.PlaystationRegisterU(0) << length) << start
	accumulator &= ^mask

	// Align the current range to where we want to insert it
	current <<= start

	return accumulator | current
}

func FillSignBits(imm arch.PlaystationRegisterU, start uint8) arch.PlaystationRegisterU {
	// Check if immediate's sign bit is set.
	if imm>>start&1 == 1 {
		// Create bitmask with (32 - start) 1s and shift it to the sign bit's position.
		mask := arch.PlaystationRegisterU(math.Pow(2, float64(32-start)) - 1)
		mask <<= start
		return imm | mask
	}

	return imm
}
