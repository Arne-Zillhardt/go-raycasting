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
		g.player.Move(FORWARD)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.Move(BACKWARD)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.Move(LEFT)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.Move(RIGHT)
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

				vector.StrokeLine(imageToDraw, float32(x), float32(y), 0, 0, 1, colorToDraw, false)
			}
			imageToDraw.Set(x, y, colorToDraw)
		}
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
