//go:build sandy32

package arch

const SandyPlatform = "sandy32"

type PlaystationRegister int32
type PlaystationRegisterU uint32

type PlaystationInstruction uint32

type PlaystationWordDouble int64
type PlaystationWordDoubleU uint64
type PlaystationWord int32
type PlaystationWordU uint32
type PlaystationWordHalf int16
type PlaystationWordHalfU uint16
type PlaystationWordByte int8
type PlaystationWordByteU uint8

var PlaystationBootRomVector = PlaystationRegisterU(0xBFC00000)
