package gpu

import (
	"fmt"

	"github.com/LamkasDev/sandy/cmd/common/arch"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type PlaystationRenderer struct {
	Window   *sdl.Window
	Surface  *sdl.Surface
	Renderer *sdl.Renderer
	Font     *ttf.Font
	Pitch    int
}

func NewRenderer() (PlaystationRenderer, error) {
	renderer := PlaystationRenderer{}
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return renderer, fmt.Errorf("failed to initialize SDL: %w", err)
	}

	w, h := int32(800), int32(600)
	renderer.Window, err = sdl.CreateWindow(fmt.Sprintf("Sandy (%s)", arch.SandyPlatform), sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		return renderer, fmt.Errorf("failed to create window: %w", err)
	}
	icon, err := img.Load("../../resources/icons/sandy.png")
	if err != nil {
		return renderer, fmt.Errorf("failed to load icon: %w", err)
	}
	renderer.Window.SetIcon(icon)
	renderer.Surface, err = renderer.Window.GetSurface()
	if err != nil {
		return renderer, fmt.Errorf("failed to get surface: %w", err)
	}
	renderer.Renderer, err = sdl.CreateRenderer(renderer.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return renderer, fmt.Errorf("failed to create renderer: %w", err)
	}

	err = ttf.Init()
	if err != nil {
		return renderer, fmt.Errorf("failed to initialize TTF: %w", err)
	}
	renderer.Font, err = ttf.OpenFont("../../resources/fonts/Roboto-Regular.ttf", 32)
	if err != nil {
		return renderer, fmt.Errorf("failed to load font: %w", err)
	}

	return renderer, nil
}

func CleanRenderer(renderer *PlaystationRenderer) error {
	err := renderer.Renderer.Destroy()
	if err != nil {
		return fmt.Errorf("failed to destroy renderer: %w", err)
	}
	err = renderer.Window.Destroy()
	if err != nil {
		return fmt.Errorf("failed to destroy window: %w", err)
	}
	renderer.Font.Close()

	sdl.Quit()
	return nil
}
