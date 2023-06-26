package cpu

import "github.com/LamkasDev/sandy/cmd/common/arch"

type PlaystationCPULoad struct {
	Register PlaystationRegisterId
	Value    arch.PlaystationRegisterU
}
