import (
	"github.com/veandco/go-sdl2/sdl"
)

err := sdl.Init(sdl.INIT_EVERYTHING)
if err != nil {
	return fmt.Errorf("Initializing SDL: %s\n", err.Error())
}
defer sdl.Quit()

window, err := sdl.CreateWindow("Flappy Gopher",
	sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 
	800, 600, sdl.WINDOW_OPENGL)
if err != nil {
	return fmt.Errorf("Initializing window: %s\n", err.Error())
}
defer window.Destroy()

renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
if err != nil {
	return fmt.Errorf("Initializing Renderer: %s\n", err.Error())
}
defer renderer.Destroy()