package instructions

import "github.com/LamkasDev/sandy/cmd/common/instruction"

// Arithmetic.
var PlaystationInstructionADD = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x20, Format: instruction.PlaystationInstructionFormatR, Name: "ADD",
}
var PlaystationInstructionADDU = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x21, Format: instruction.PlaystationInstructionFormatR, Name: "ADDU",
}
var PlaystationInstructionSUB = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x22, Format: instruction.PlaystationInstructionFormatR, Name: "SUB",
}
var PlaystationInstructionSUBU = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x23, Format: instruction.PlaystationInstructionFormatR, Name: "SUBU",
}
var PlaystationInstructionMULT = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x18, Format: instruction.PlaystationInstructionFormatR, Name: "MULT",
}
var PlaystationInstructionMULTU = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x19, Format: instruction.PlaystationInstructionFormatR, Name: "MULTU",
}
var PlaystationInstructionDIV = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x1A, Format: instruction.PlaystationInstructionFormatR, Name: "DIV",
}
var PlaystationInstructionDIVU = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x1B, Format: instruction.PlaystationInstructionFormatR, Name: "DIVU",
}

var PlaystationInstructionADDI = instruction.PlaystationInstructionDefinition{
	Code: 0b001000, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "ADDI",
}
var PlaystationInstructionADDIU = instruction.PlaystationInstructionDefinition{
	Code: 0b001001, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "ADDIU",
}
var PlaystationInstructionLUI = instruction.PlaystationInstructionDefinition{
	Code: 0b001111, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "LUI",
}

// Logical.
var PlaystationInstructionSLL = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x0, Format: instruction.PlaystationInstructionFormatR, Name: "SLL",
}
var PlaystationInstructionSRL = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x2, Format: instruction.PlaystationInstructionFormatR, Name: "SRL",
}
var PlaystationInstructionSRA = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x3, Format: instruction.PlaystationInstructionFormatR, Name: "SRA",
}
var PlaystationInstructionSLLV = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x4, Format: instruction.PlaystationInstructionFormatR, Name: "SLLV",
}
var PlaystationInstructionSRLV = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x6, Format: instruction.PlaystationInstructionFormatR, Name: "SRLV",
}
var PlaystationInstructionSRAV = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x7, Format: instruction.PlaystationInstructionFormatR, Name: "SRAV",
}

var PlaystationInstructionAND = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x24, Format: instruction.PlaystationInstructionFormatR, Name: "AND",
}
var PlaystationInstructionOR = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x25, Format: instruction.PlaystationInstructionFormatR, Name: "OR",
}
var PlaystationInstructionXOR = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x26, Format: instruction.PlaystationInstructionFormatR, Name: "XOR",
}
var PlaystationInstructionNOR = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x27, Format: instruction.PlaystationInstructionFormatR, Name: "NOR",
}

var PlaystationInstructionANDI = instruction.PlaystationInstructionDefinition{
	Code: 0b001100, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "ANDI",
}
var PlaystationInstructionORI = instruction.PlaystationInstructionDefinition{
	Code: 0b001101, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "ORI",
}
var PlaystationInstructionXORI = instruction.PlaystationInstructionDefinition{
	Code: 0b001110, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "XORI",
}

var PlaystationInstructionSLT = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x2A, Format: instruction.PlaystationInstructionFormatR, Name: "SLT",
}
var PlaystationInstructionSLTU = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x2B, Format: instruction.PlaystationInstructionFormatR, Name: "SLTU",
}

var PlaystationInstructionSLTI = instruction.PlaystationInstructionDefinition{
	Code: 0b001010, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "SLTI",
}
var PlaystationInstructionSLTIU = instruction.PlaystationInstructionDefinition{
	Code: 0b001011, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "SLTIU",
}

var PlaystationInstructionMFLO = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x12, Format: instruction.PlaystationInstructionFormatR, Name: "MFLO",
}
var PlaystationInstructionMFHI = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x10, Format: instruction.PlaystationInstructionFormatR, Name: "MFHI",
}

// Load / Store.
var PlaystationInstructionLW = instruction.PlaystationInstructionDefinition{
	Code: 0b100011, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "LW",
}
var PlaystationInstructionLB = instruction.PlaystationInstructionDefinition{
	Code: 0b100000, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "LB",
}
var PlaystationInstructionLBU = instruction.PlaystationInstructionDefinition{
	Code: 0b100100, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "LBU",
}

var PlaystationInstructionSW = instruction.PlaystationInstructionDefinition{
	Code: 0b101011, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "SW",
}
var PlaystationInstructionSH = instruction.PlaystationInstructionDefinition{
	Code: 0b101001, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "SW",
}
var PlaystationInstructionSB = instruction.PlaystationInstructionDefinition{
	Code: 0b101000, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "SB",
}

// Branching.
var PlaystationInstructionJ = instruction.PlaystationInstructionDefinition{
	Code: 0b000010, Argument: 0x0, Format: instruction.PlaystationInstructionFormatJ, Name: "J",
}
var PlaystationInstructionBNE = instruction.PlaystationInstructionDefinition{
	Code: 0b000101, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "BNE",
}
var PlaystationInstructionJAL = instruction.PlaystationInstructionDefinition{
	Code: 0b000011, Argument: 0x0, Format: instruction.PlaystationInstructionFormatJ, Name: "JAL",
}
var PlaystationInstructionJR = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x8, Format: instruction.PlaystationInstructionFormatR, Name: "JR",
}
var PlaystationInstructionBEQ = instruction.PlaystationInstructionDefinition{
	Code: 0b000100, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "BEQ",
}
var PlaystationInstructionBGTZ = instruction.PlaystationInstructionDefinition{
	Code: 0b000111, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "BGTZ",
}
var PlaystationInstructionBLEZ = instruction.PlaystationInstructionDefinition{
	Code: 0b000110, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "BLEZ",
}
var PlaystationInstructionJALR = instruction.PlaystationInstructionDefinition{
	Code: 0b000000, Argument: 0x9, Format: instruction.PlaystationInstructionFormatR, Name: "JALR",
}
var PlaystationInstructionBZ = instruction.PlaystationInstructionDefinition{
	Code: 0b000001, Argument: 0x0, Format: instruction.PlaystationInstructionFormatI, Name: "BZ",
}

// Coprocessor 0.
var PlaystationInstructionMTC0 = instruction.PlaystationInstructionDefinition{
	Code: 0b010000, Argument: 0x4, Format: instruction.PlaystationInstructionFormatCP, Name: "MTC0",
}
var PlaystationInstructionMFC0 = instruction.PlaystationInstructionDefinition{
	Code: 0b010000, Argument: 0x0, Format: instruction.PlaystationInstructionFormatCP, Name: "MFC0",
}
