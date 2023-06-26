package util

import "github.com/LamkasDev/sandy/cmd/common/arch"

func SignExtendRegisterU(value arch.PlaystationRegisterU) arch.PlaystationRegisterU {
	return arch.PlaystationRegisterU(int16(value))
}

func SignExtendByte(value arch.PlaystationWordByte) arch.PlaystationRegisterU {
	return arch.PlaystationRegisterU(int8(value))
}
