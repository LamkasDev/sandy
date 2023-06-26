package cpu

import "github.com/LamkasDev/sandy/cmd/common/arch"

type PlaystationCP0 struct {
	Snapshot  [32]arch.PlaystationRegisterU
	Registers [32]arch.PlaystationRegisterU
}

func NewPlaystationCP0() PlaystationCP0 {
	return PlaystationCP0{}
}
