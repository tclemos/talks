import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (g *Game) checkPollEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			g.quit()
			break
		}
	}
}

func (g *Game) checkKeys() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_RETURN] == 1 {
		g.restart()
	} else if keys[sdl.SCANCODE_ESCAPE] == 1 {
		g.quit()
	}
}

func (g *Game) drawBackGround() {
	background := &sdl.Rect{0, 0, screenWidth, screenHeight}
	g.renderer.SetDrawColor(255, 0, 0, 255) // RGBA
	g.renderer.FillRect(background)
}

func (g *Game) drawTexture() {
	image, _ := img.Load("../game/sprites/cloud.png")
	tx, _ := g.renderer.CreateTextureFromSurface(image)
	src := &sdl.Rect{0, 0, 100, 100}
	dst := &sdl.Rect{50, 50, 200, 200}
	g.renderer.Copy(tx, src, dst)
}