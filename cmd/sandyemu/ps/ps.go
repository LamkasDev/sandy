package ps

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime/pprof"
	"time"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/LamkasDev/sandy/cmd/common/cpu"
	"github.com/LamkasDev/sandy/cmd/common/instruction"
	"github.com/LamkasDev/sandy/cmd/common/instructions"
	"github.com/LamkasDev/sandy/cmd/sandyemu/config"
	"github.com/LamkasDev/sandy/cmd/sandyemu/decoder"
	"github.com/LamkasDev/sandy/cmd/sandyemu/gpu"
	"github.com/LamkasDev/sandy/cmd/sandyemu/instruction_runners"
	"github.com/LamkasDev/sandy/cmd/sandyemu/internal"
	"github.com/LamkasDev/sandy/cmd/sandyemu/memory"
	"github.com/jwalton/gchalk"
	"github.com/rs/zerolog/log"
)

type Playstation struct {
	Config   config.PlaystationConfig
	Internal internal.PlaystationInternal
	Table    instructions.PlaystationInstructionTable
}

func NewPlaystation() Playstation {
	// speed := 1.0
	cps := Playstation{
		Config:   config.NewPlaystationConfig(),
		Internal: internal.NewInternal(),
		Table:    instruction_runners.NewDecoderInstructionTable(),
	}

	u, err := user.Current()
	if err != nil {
		log.Error().Msgf("failed to get user: %v", err)
		panic(err)
	}
	if err = os.MkdirAll(filepath.Join(u.HomeDir, "Documents", "sandy", "config"), 0755); err != nil {
		log.Error().Msgf("failed to create data directory: %v", err)
		panic(err)
	}
	if err = config.ReadConfig(&cps.Config); err != nil {
		log.Error().Msgf("failed to read config: %v", err)
		panic(err)
	}

	return cps
}

func LoadPlaystation(cps *Playstation, bios string) {
	biosFile, err := os.ReadFile(bios)
	if err != nil {
		log.Error().Msgf("failed to read bios file: %v", err)
		panic(err)
	}

	memory.WriteMemory(&cps.Internal.Memory, arch.PlaystationBootRomVector, biosFile)
}

func RunPlaystation(cps *Playstation) {
	log.Info().Msgf("hi from %s :3", gchalk.Red(arch.SandyPlatform))
	defer gpu.CleanGPU(&cps.Internal.GPU)
	defer ClosePlaystation(cps)

	go func() {
		defer ClosePlaystation(cps)

		// cycles := cps.Config.CPU.Speed / 1000
		for cps.Internal.Executing {
			// for i := uint32(0); i < cycles && cps.Internal.Executing; i++ {
			RunPlaystationCycle(cps)
			cps.Internal.CPU.TotalCycles++
			// }
			// time.Sleep(time.Millisecond)
		}
	}()
	go func() {
		defer ClosePlaystation(cps)

		for cps.Internal.Executing {
			if cps.Internal.CPU.Cache == nil {
				CacheInstruction(cps)
			}
		}
	}()
	for cps.Internal.Running {
		internal.HandleEventsGPU(&cps.Internal)
		gpu.RenderGPU(&cps.Internal.GPU)
		time.Sleep(time.Millisecond)
	}
}

var ClosedAlready = false

func ClosePlaystation(cps *Playstation) {
	if ClosedAlready {
		return
	}
	ClosedAlready = true

	u, _ := user.Current()
	dumpFile, err := os.Create(path.Join(u.HomeDir, "Documents", "sandy", "ram.dump"))
	if err != nil {
		log.Error().Msgf("failed to open dump file: %v", err)
		panic(err)
	}
	dumpFile.Write(cps.Internal.Memory.Sections[len(cps.Internal.Memory.Sections)-1].Data[0:67108864])
	dumpFile.Close()

	log.Info().Msgf("cycles: %s;",
		gchalk.Red(fmt.Sprint(cps.Internal.CPU.TotalCycles)),
	)
	log.Info().Msg("bay! :33")
	pprof.StopCPUProfile()
	os.Exit(0)
}

func CacheInstruction(cps *Playstation) {
	rawNextInstruction := arch.PlaystationInstruction(memory.ReadMemoryWord(&cps.Internal.Memory, cps.Internal.CPU.CacheCounter))
	nextInstruction, err := decoder.Decode(cps.Table, rawNextInstruction, cps.Internal.CPU.CacheCounter)
	if err != nil {
		log.Error().Msgf("failed to decode instruction: %v (c: %s)", err, gchalk.Red(fmt.Sprintf("%#08x", cps.Internal.CPU.CacheCounter)))
		panic(err)
	}
	cps.Internal.CPU.Cache = &nextInstruction
}

func RunPlaystationCycle(cps *Playstation) {
	for cps.Internal.CPU.Cache == nil {
		continue
	}
	currentInstruction := *cps.Internal.CPU.Cache
	cps.Internal.CPU.CacheCounter = cps.Internal.CPU.Registers.Counter
	cps.Internal.CPU.Cache = nil
	cps.Internal.CPU.Registers.Counter += 4

	if cps.Internal.CPU.Load != nil {
		cps.Internal.CPU.Registers.NextGeneral[cps.Internal.CPU.Load.Register] = cps.Internal.CPU.Load.Value
		cps.Internal.CPU.Load = nil
	}
	cps.Internal.CPU.Registers.General[cpu.PlaystationRegisterZero] = 0

	if arch.SandyDebug {
		internal.PrintInternalInstruction(currentInstruction)
	}
	currentInstruction.Process.(func(*internal.PlaystationInternal, instruction.PlaystationInstruction))(&cps.Internal, currentInstruction)

	cps.Internal.CPU.Registers.General = cps.Internal.CPU.Registers.NextGeneral
}
