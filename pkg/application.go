package pkg

import (
	"time"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const w int = 640
const h int = 480
const planeX float64 = 0
const planeY float64 = 0.66


type Application struct {
	gameMap [][]uint
	player  player
}

var oldTime time.Time

func (a *Application) Update() error {
	currentTime := time.Now()
	var deltaTime float64 = float64(currentTime.UnixNano() - oldTime.UnixNano())/1000000000

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		a.player.move(FORWARD, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		a.player.move(BACKWARD, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		a.player.move(LEFT, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		a.player.move(RIGHT, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		a.player.rotateRight(deltaTime)
	}

	oldTime = time.Now()

	return nil
}

func (a *Application) Draw(screen *ebiten.Image) {
	imageToDraw := ebiten.NewImage(w, h)
	drawFloorAndCeiling(imageToDraw)

	for x := range(w) {
		var cameraX float64 = 2 * float64(x) / float64(w) - 1
		var rayDirX float64 = a.player.dirX + planeX * cameraX
		var rayDirY float64 = a.player.dirY + planeY * cameraX

		distanceToWall, side := calculateRayLength(a.player, rayDirX, rayDirY, a.gameMap)

		lineHeight := (int)(float64(h) / distanceToWall)
		drawStart := -lineHeight / 2 + h / 2
		if drawStart < 0 {
			drawStart = 0
		}
		drawEnd := lineHeight / 2 + h / 2
		if drawEnd >= h {
			drawEnd = h-1
		}

		colorToDraw := color.RGBA{0xff, 0xff, 0xff, 0xff}
		if side == 1 {
			colorToDraw = color.RGBA{0x7f, 0x7f, 0x7f, 0xff}
		}

		vector.StrokeLine(imageToDraw, float32(x), float32(drawStart), float32(x), float32(drawEnd), 1, colorToDraw, false)
	}

	screen.DrawImage(imageToDraw, nil)
}

func drawFloorAndCeiling(screen *ebiten.Image) {
	for y := range(h) {
		colorToDraw := color.RGBA{0x70, 0x70, 0x70, 0xff}

		if y <= h/2 {
			colorToDraw = color.RGBA{0x38, 0x38, 0x38, 0xff}
		}

		vector.StrokeLine(screen, 0,float32(y), float32(w), float32(y),  1, colorToDraw, false)
	}
}

func (a *Application) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w, h
}

func (a *Application) Init() {
	a.gameMap = getGameMap(24, 24)

	a.player = createPlayer(a.gameMap)
}
