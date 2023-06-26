package instruction

const PlaystationInstructionFormatR = uint8(0x00)
const PlaystationInstructionFormatI = uint8(0x01)
const PlaystationInstructionFormatJ = uint8(0x02)
const PlaystationInstructionFormatCP = uint8(0x03)

type PlaystationInstructionFormat struct {
	Code        PlaystationInstructionRange
	Argument    PlaystationInstructionRange
	Source      PlaystationInstructionRange
	Target      PlaystationInstructionRange
	Destination PlaystationInstructionRange
	Immediate   PlaystationInstructionRange
}

var PlaystationInstructionFormats = map[uint8]PlaystationInstructionFormat{
	PlaystationInstructionFormatR: {
		Argument:    PlaystationInstructionRange{Start: 0, Length: 6},
		Immediate:   PlaystationInstructionRange{Start: 6, Length: 5},
		Destination: PlaystationInstructionRange{Start: 11, Length: 5},
		Target:      PlaystationInstructionRange{Start: 16, Length: 5},
		Source:      PlaystationInstructionRange{Start: 21, Length: 5},
		Code:        PlaystationInstructionRange{Start: 26, Length: PlaystationInstructionCodeLength},
	},
	PlaystationInstructionFormatI: {
		Immediate: PlaystationInstructionRange{Start: 0, Length: 16},
		Target:    PlaystationInstructionRange{Start: 16, Length: 5},
		Source:    PlaystationInstructionRange{Start: 21, Length: 5},
		Code:      PlaystationInstructionRange{Start: 26, Length: PlaystationInstructionCodeLength},
	},
	PlaystationInstructionFormatJ: {
		Immediate: PlaystationInstructionRange{Start: 0, Length: 26},
		Code:      PlaystationInstructionRange{Start: 26, Length: PlaystationInstructionCodeLength},
	},
	PlaystationInstructionFormatCP: {
		Immediate:   PlaystationInstructionRange{Start: 0, Length: 11},
		Destination: PlaystationInstructionRange{Start: 11, Length: 5},
		Target:      PlaystationInstructionRange{Start: 16, Length: 5},
		Argument:    PlaystationInstructionRange{Start: 21, Length: 5},
		Code:        PlaystationInstructionRange{Start: 26, Length: PlaystationInstructionCodeLength},
	},
}
