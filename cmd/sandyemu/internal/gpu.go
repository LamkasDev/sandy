package internal

import (
	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/sdl"
)

func HandleEventsGPU(cinternal *PlaystationInternal) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			cinternal.Executing = false
			cinternal.Running = false
		case *sdl.WindowEvent:
			switch event.(*sdl.WindowEvent).Event {
			case sdl.WINDOWEVENT_RESIZED:
				var err error
				if cinternal.GPU.Renderer.Surface, err = cinternal.GPU.Renderer.Window.GetSurface(); err != nil {
					log.Error().Msgf("failed to get surface: %v", err)
					panic(err)
				}
			}
		case *sdl.KeyboardEvent:
			if t.Type == sdl.KEYDOWN {
				switch t.Keysym.Scancode {
				case sdl.SCANCODE_D:
					arch.SandyDebug = !arch.SandyDebug
				}
			}
		}
	}
}
