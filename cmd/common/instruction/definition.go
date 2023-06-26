package instruction

const PlaystationInstructionCodeLength = 6

type PlaystationInstructionRange struct {
	Start  uint8
	Length uint8
}

type PlaystationInstructionDefinition struct {
	Code     uint8
	Argument uint8
	Format   uint8
	Name     string
}
