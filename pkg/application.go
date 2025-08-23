package pkg

import (
	"github.com/hajimehoshi/ebiten/v2"
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
}

func (a *Application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (a *Application) Init() {
	a.gameMap = getGameMap(24, 24)
	a.player = createPlayer(a.gameMap)
}
