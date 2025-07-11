package game

import (
	"image/color"
	"log"
	"time"

	//"github.com/arne-zillhardt/raycasting/pkg/raycasting"
	"github.com/arne-zillhardt/raycasting/pkg/raycasting"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game interface {
	Update() error
	Draw(*ebiten.Image)
	Layout(int, int) (int, int)
	PlayerPosition() (int, int)
	IsWall(uint32, uint32) bool
	Size() int
}
type game struct {
	playingField gameMap
	player       player
}

var w int = 320
var h int = 200
var oldTime time.Time

func (g game) Update() error {
	currentTime := time.Now()
	var deltaTime float64 = float64(currentTime.UnixNano() - oldTime.UnixNano())/1000000000

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.move(FORWARD, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.move(BACKWARD, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.move(LEFT, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.move(RIGHT, deltaTime)
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.player.changeView(INCREASE, deltaTime)
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		g.player.changeView(DECREASE, deltaTime)
	}

	oldTime = time.Now()
	return nil
}

func (g game) Draw(screen *ebiten.Image) {
	imageToDraw := ebiten.NewImage(w, h)

	colorToDraw := color.RGBA{0x3C, 0x3C, 0xE0, 0xff}
	for y := range h {
		for x := range w {
			colorToDraw = color.RGBA{0x70, 0x70, 0x70, 0xff}

			if y <= h/2 {
				colorToDraw = color.RGBA{0x38, 0x38, 0x38, 0xff}
			}

			imageToDraw.Set(x, y, colorToDraw)
		}
	}

	playerPositionX, playerPositionY := g.player.position()
	startFrom := -(float64(w)/2)
	for x := range w {
		colorToDraw = color.RGBA{0xff, 0xff, 0xff, 0xff}
		vectorPositionX, vectorPositionY := g.player.calculateLinePosition(startFrom)
		positionX := float64(playerPositionX) * vectorPositionX
		positionY := float64(playerPositionY) * vectorPositionY

  	rayLength := raycasting.CalculateRay(float64(playerPositionX), float64(playerPositionY), positionX, positionY, g.playingField.playMap)
		vector.StrokeLine(imageToDraw, float32(x), float32(0), float32(x), float32(rayLength), 1, colorToDraw, false)
		startFrom++

		log.Println(x)
  }

	fieldSize := g.playingField.size()
	for y := range fieldSize {
		for x := range fieldSize {
			colorToDraw = color.RGBA{0x3C, 0x3C, 0xE0, 0xff}
			if !g.playingField.isWall(uint32(x), uint32(y)) {
				colorToDraw = color.RGBA{0xB5, 0x9B, 0x41, 0xff}
			}
			
			if x == playerPositionX && y == playerPositionY {
				colorToDraw = color.RGBA{0xff, 0x00, 0x00, 0xff}
			}

			imageToDraw.Set(x, y, colorToDraw)
		}
	}

	screen.DrawImage(imageToDraw, nil)
}

func (g game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w, h
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

func (g game) Size() int {
	return g.playingField.size()
}
