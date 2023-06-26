package instruction_runners

import "github.com/LamkasDev/sandy/cmd/common/instructions"

func NewDecoderInstructionTable() instructions.PlaystationInstructionTable {
	table := instructions.NewInstructionTable()

	// Arithmetic.
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionLUI.Code, instructions.PlaystationInstructionLUI.Argument, ProcessLUI)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionADDIU.Code, instructions.PlaystationInstructionADDIU.Argument, ProcessADDIU)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionADDI.Code, instructions.PlaystationInstructionADDI.Argument, ProcessADDI)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionADDU.Code, instructions.PlaystationInstructionADDU.Argument, ProcessADDU)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionADD.Code, instructions.PlaystationInstructionADD.Argument, ProcessADD)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSUBU.Code, instructions.PlaystationInstructionSUBU.Argument, ProcessSUBU)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionDIV.Code, instructions.PlaystationInstructionDIV.Argument, ProcessDIV)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionDIVU.Code, instructions.PlaystationInstructionDIVU.Argument, ProcessDIVU)

	// Logical.
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionORI.Code, instructions.PlaystationInstructionORI.Argument, ProcessORI)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSLL.Code, instructions.PlaystationInstructionSLL.Argument, ProcessSLL)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionOR.Code, instructions.PlaystationInstructionOR.Argument, ProcessOR)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSLTU.Code, instructions.PlaystationInstructionSLTU.Argument, ProcessSLTU)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionANDI.Code, instructions.PlaystationInstructionANDI.Argument, ProcessANDI)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionAND.Code, instructions.PlaystationInstructionAND.Argument, ProcessAND)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSLTI.Code, instructions.PlaystationInstructionSLTI.Argument, ProcessSLTI)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSRA.Code, instructions.PlaystationInstructionSRA.Argument, ProcessSRA)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionMFLO.Code, instructions.PlaystationInstructionMFLO.Argument, ProcessMFLO)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSRL.Code, instructions.PlaystationInstructionSRL.Argument, ProcessSRL)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSLTIU.Code, instructions.PlaystationInstructionSLTIU.Argument, ProcessSLTIU)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionMFHI.Code, instructions.PlaystationInstructionMFHI.Argument, ProcessMFHI)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSLT.Code, instructions.PlaystationInstructionSLT.Argument, ProcessSLT)

	// Load / Store.
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionLW.Code, instructions.PlaystationInstructionLW.Argument, ProcessLW)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSW.Code, instructions.PlaystationInstructionSW.Argument, ProcessSW)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSH.Code, instructions.PlaystationInstructionSH.Argument, ProcessSH)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionSB.Code, instructions.PlaystationInstructionSB.Argument, ProcessSB)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionLB.Code, instructions.PlaystationInstructionLB.Argument, ProcessLB)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionLBU.Code, instructions.PlaystationInstructionLBU.Argument, ProcessLBU)

	// Branching.
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionJ.Code, instructions.PlaystationInstructionJ.Argument, ProcessJ)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionBNE.Code, instructions.PlaystationInstructionBNE.Argument, ProcessBNE)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionJAL.Code, instructions.PlaystationInstructionJAL.Argument, ProcessJAL)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionJR.Code, instructions.PlaystationInstructionJR.Argument, ProcessJR)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionBEQ.Code, instructions.PlaystationInstructionBEQ.Argument, ProcessBEQ)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionBGTZ.Code, instructions.PlaystationInstructionBGTZ.Argument, ProcessBGTZ)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionBLEZ.Code, instructions.PlaystationInstructionBLEZ.Argument, ProcessBLEZ)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionJALR.Code, instructions.PlaystationInstructionJALR.Argument, ProcessJALR)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionBZ.Code, instructions.PlaystationInstructionBZ.Argument, ProcessBZ)

	// CP0.
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionMTC0.Code, instructions.PlaystationInstructionMTC0.Argument, ProcessMTC0)
	instructions.DecorateInstructionTableEntry(table, instructions.PlaystationInstructionMFC0.Code, instructions.PlaystationInstructionMFC0.Argument, ProcessMFC0)

	return table
}
