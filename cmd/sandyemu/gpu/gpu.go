package gpu

import (
	"github.com/rs/zerolog/log"
	"github.com/veandco/go-sdl2/sdl"
)

type PlaystationGPU struct {
	Ticks    uint32
	Renderer PlaystationRenderer
}

func NewPlaystationGPU() PlaystationGPU {
	renderer, err := NewRenderer()
	if err != nil {
		log.Error().Msgf("failed to initialize renderer: %v", err)
		panic(err)
	}

	return PlaystationGPU{
		Renderer: renderer,
	}
}

func RenderGPU(cgpu *PlaystationGPU) {
	if cgpu.Ticks > 1000/PlaystationGPUFps {
		if err := cgpu.Renderer.Surface.FillRect(nil, 0); err != nil {
			log.Error().Msgf("failed to fill rect: %v", err)
			panic(err)
		}
		text, err := cgpu.Renderer.Font.RenderUTF8Solid("Playstation", sdl.Color{B: 255, A: 255})
		if err != nil {
			log.Error().Msgf("failed to render text: %v", err)
			panic(err)
		}
		if err := text.Blit(nil, cgpu.Renderer.Surface, &sdl.Rect{X: 10, Y: 10}); err != nil {
			log.Error().Msgf("failed to blit text: %v", err)
			panic(err)
		}
		if err := cgpu.Renderer.Window.UpdateSurface(); err != nil {
			log.Error().Msgf("failed to update surface: %v", err)
			panic(err)
		}
		cgpu.Ticks = 0
	}
}

func CleanGPU(cgpu *PlaystationGPU) error {
	return CleanRenderer(&cgpu.Renderer)
}
