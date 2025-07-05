package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game interface {
	Update() error
	Draw(*ebiten.Image)
	Layout(int, int) (int, int)
	PlayerPosition() (int, int)
	IsWall(uint32, uint32) bool
}
type game struct {
	playingField gameMap
	player       player
}

func (g game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.move(FORWARD)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.move(BACKWARD)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.move(LEFT)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.move(RIGHT)
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.player.changeView(INCREASE)
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.player.changeView(DECREASE)
	}

	return nil
}

func (g game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	imageToDraw := ebiten.NewImage(320, 200)

	for y := range 200 {
		for x := range 320 {
			colorToDraw := color.RGBA{0x70, 0x70, 0x70, 0xff}

			if y <= 200/2 {
				colorToDraw = color.RGBA{0x38, 0x38, 0x38, 0xff}
			}

			if *g.player.positonX == int32(x) && *g.player.positonY == int32(y) {
				colorToDraw = color.RGBA{0xff, 0xff, 0xff, 0xff}

			}
			imageToDraw.Set(x, y, colorToDraw)
		}

		colorToDraw := color.RGBA{0xff, 0xff, 0xff, 0xff}
		lineX, lineY := g.player.calculateLinePosition()
		playerX, playerY := g.player.position()
		vector.StrokeLine(imageToDraw, float32(playerX), float32(playerY), float32(lineX), float32(lineY), 1, colorToDraw, false)
	}

	screen.DrawImage(imageToDraw, nil)
}

func (g game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 200
}

func GetInstance() Game {
	log.Println("Created game")
	return game{
		player:       initPlayer(),
		playingField: initGameMap(),
	}
}

func (g game) PlayerPosition() (int, int) {
	return g.player.position()
}

func (g game) IsWall(positionX uint32, positionY uint32) bool {
	return g.playingField.isWall(positionX, positionY)
}
