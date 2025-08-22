package pkg

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const h int = 320
const w int = 240

type Application struct {
	gameMap [][]uint
	player  player
}

func (a *Application) Update() error {
	return nil
}

func (a *Application) Draw(screen *ebiten.Image) {
	for y := range h {
		for x := range w {
			screen.Set(x, y, color.RGBA{142, 142, 142, 1})

			if y >= h/2 {
				screen.Set(x, y, color.RGBA{63, 70, 71, 1})
			}
		}
	}
	for x := range h {
		wallDistance, side := calculateRayLength(a.player.positionX, a.player.positionY, 1, 0.66, a.gameMap)

		colorToDraw := color.RGBA{255, 255, 255, 1}
		if side == 1 {
			colorToDraw = color.RGBA{127, 127, 127, 1}
		}

		var lineHeight int = int(h / int(wallDistance))

		var drawStart int = -lineHeight/2 + h/2
		if drawStart < 0 {
			drawStart = 0
		}
		var drawEnd int = lineHeight/2 + h/2
		if drawEnd >= h {
			drawEnd = h - 1
		}

		vector.StrokeLine(screen, float32(x), float32(drawStart), float32(x), float32(drawEnd), 1.0, colorToDraw, false)
	}
}

func (a *Application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (a *Application) Init() {
	a.gameMap = getGameMap(24, 24)
	a.player = createPlayer(a.gameMap)
}
