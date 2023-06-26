package instructions

import (
	"github.com/LamkasDev/sandy/cmd/common/instruction"
)

type PlaystationInstructionTableEntry struct {
	Instruction instruction.PlaystationInstructionDefinition
	Process     interface{}
}
type PlaystationInstructionTableSubtable struct {
	Format   uint8
	Subtable map[uint8]PlaystationInstructionTableEntry
}
type PlaystationInstructionTable map[uint8]PlaystationInstructionTableSubtable

func SetupInstructionTableEntry(table PlaystationInstructionTable, cinstruction instruction.PlaystationInstructionDefinition, process interface{}) {
	_, ok := table[cinstruction.Code]
	if !ok {
		table[cinstruction.Code] = PlaystationInstructionTableSubtable{
			Format:   cinstruction.Format,
			Subtable: map[uint8]PlaystationInstructionTableEntry{},
		}
	}

	table[cinstruction.Code].Subtable[cinstruction.Argument] = PlaystationInstructionTableEntry{
		Instruction: cinstruction,
		Process:     process,
	}
}

func DecorateInstructionTableEntry(table PlaystationInstructionTable, code uint8, argument uint8, process interface{}) {
	e := table[code].Subtable[argument]
	e.Process = process
	table[code].Subtable[argument] = e
}

func NewInstructionTable() PlaystationInstructionTable {
	table := PlaystationInstructionTable{}

	// Arithmetic.
	SetupInstructionTableEntry(table, PlaystationInstructionLUI, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionADDIU, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionADDI, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionADDU, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionADD, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSUBU, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionDIV, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionDIVU, nil)

	// Logical.
	SetupInstructionTableEntry(table, PlaystationInstructionORI, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSLL, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionOR, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSLTU, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionANDI, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionAND, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSLTI, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSRA, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionMFLO, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSRL, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSLTIU, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionMFHI, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSLT, nil)

	// Load / Store.
	SetupInstructionTableEntry(table, PlaystationInstructionLW, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSW, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSH, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionSB, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionLB, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionLBU, nil)

	// Branching.
	SetupInstructionTableEntry(table, PlaystationInstructionJ, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionBNE, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionJAL, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionJR, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionBEQ, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionBGTZ, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionBLEZ, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionBZ, nil)

	// CP0.
	SetupInstructionTableEntry(table, PlaystationInstructionMTC0, nil)
	SetupInstructionTableEntry(table, PlaystationInstructionMFC0, nil)

	return table
}
