package cpu

type PlaystationRegisterId uint8

var PlaystationRegisterNames = map[PlaystationRegisterId]string{
	0x00: "zero",
	0x01: "at",
	0x02: "v0",
	0x03: "v1",
	0x04: "a0",
	0x05: "a1",
	0x06: "a2",
	0x07: "a3",
	0x08: "t0",
	0x09: "t1",
	0x0a: "t2",
	0x0b: "t3",
	0x0c: "t4",
	0x0d: "t5",
	0x0e: "t6",
	0x0f: "t7",
	0x10: "s0",
	0x11: "s1",
	0x12: "s2",
	0x13: "s3",
	0x14: "s4",
	0x15: "s5",
	0x16: "s6",
	0x17: "s7",
	0x18: "t8",
	0x19: "t9",
	0x1a: "k0",
	0x1b: "k1",
	0x1c: "gp",
	0x1d: "sp",
	0x1e: "fp",
	0x1f: "ra",
}

const PlaystationRegisterZero = PlaystationRegisterId(0x00)
const PlaystationRegisterAssemblerTemporary = PlaystationRegisterId(0x01)
const PlaystationRegisterReturnZero = PlaystationRegisterId(0x02)
const PlaystationRegisterReturnOne = PlaystationRegisterId(0x03)
const PlaystationRegisterArgumentZero = PlaystationRegisterId(0x04)
const PlaystationRegisterArgumentOne = PlaystationRegisterId(0x05)
const PlaystationRegisterArgumentTwo = PlaystationRegisterId(0x06)
const PlaystationRegisterArgumentThree = PlaystationRegisterId(0x07)
const PlaystationRegisterTemporaryZero = PlaystationRegisterId(0x08)
const PlaystationRegisterTemporaryOne = PlaystationRegisterId(0x09)
const PlaystationRegisterTemporaryTwo = PlaystationRegisterId(0x0a)
const PlaystationRegisterTemporaryThree = PlaystationRegisterId(0x0b)
const PlaystationRegisterTemporaryFour = PlaystationRegisterId(0x0c)
const PlaystationRegisterTemporaryFive = PlaystationRegisterId(0x0d)
const PlaystationRegisterTemporarySix = PlaystationRegisterId(0x0e)
const PlaystationRegisterTemporarySeven = PlaystationRegisterId(0x0f)
const PlaystationRegisterSavedZero = PlaystationRegisterId(0x10)
const PlaystationRegisterSavedOne = PlaystationRegisterId(0x11)
const PlaystationRegisterSavedTwo = PlaystationRegisterId(0x12)
const PlaystationRegisterSavedThree = PlaystationRegisterId(0x13)
const PlaystationRegisterSavedFour = PlaystationRegisterId(0x14)
const PlaystationRegisterSavedFive = PlaystationRegisterId(0x15)
const PlaystationRegisterSavedSix = PlaystationRegisterId(0x16)
const PlaystationRegisterSavedSeven = PlaystationRegisterId(0x17)
const PlaystationRegisterTemporaryEight = PlaystationRegisterId(0x18)
const PlaystationRegisterTemporaryNine = PlaystationRegisterId(0x19)
const PlaystationRegisterKernelZero = PlaystationRegisterId(0x1a)
const PlaystationRegisterKernelOne = PlaystationRegisterId(0x1b)
const PlaystationRegisterGlobalPointer = PlaystationRegisterId(0x1c)
const PlaystationRegisterStackPointer = PlaystationRegisterId(0x1d)
const PlaystationRegisterFramePointer = PlaystationRegisterId(0x1e)
const PlaystationRegisterReturnAddress = PlaystationRegisterId(0x1f)
