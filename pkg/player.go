package pkg

import (
	"log"
	"math"
)

type player struct {
	movementSpeed float64
	rotationSpeed float64
	positionX float64
	positionY float64
	dirX float64
	dirY float64
	planeX float64
	planeY float64
}

type movementDirection struct {
	addToX float64
	addToY float64
}

type rotationDirection struct {
	addToX float64
	addToY float64
}

var (
	FORWARD  movementDirection = movementDirection{0, -1}
  BACKWARD movementDirection = movementDirection{0, 1}
  LEFT     movementDirection = movementDirection{-1, 0}
  RIGHT    movementDirection = movementDirection{1, 0}
)

func (p *player) move(direction movementDirection, deltaTime float64) {
	log.Println("Before moving: x = ", p.positionX, "; y = ", p.positionY)

	p.positionX += direction.addToX * (p.movementSpeed * deltaTime)
	p.positionY += direction.addToY * (p.movementSpeed * deltaTime)

	log.Println("After moving: x = ", p.positionX, "; y = ", p.positionY)
}

func (p *player) rotateRight(deltaTime float64) {
	log.Println("Before rotating: dirX = ", p.dirX, "; dirY = ", p.dirY)
	log.Println("Before rotating: planeX = ", p.planeX, "; planeY = ", p.planeY)

	rotationSpeed := p.rotationSpeed * deltaTime
	oldDirX := float64(p.dirX)
	p.dirX = p.dirX * math.Cos(-rotationSpeed) - p.dirY * math.Sin(-rotationSpeed)
	p.dirY = oldDirX * math.Sin(-rotationSpeed) + p.dirY * math.Cos(-rotationSpeed)

	oldPlaneX := float64(p.planeX)
	p.planeX = p.planeX * math.Cos(-rotationSpeed) - p.planeY * math.Sin(-rotationSpeed)
	p.planeY = oldPlaneX * math.Sin(-rotationSpeed) + p.planeY * math.Cos(-rotationSpeed)

	log.Println("After rotating: dirX = ", p.dirX, "; dirY = ", p.dirY)
	log.Println("After rotating: planeX = ", p.planeX, "; planeY = ", p.planeY)
}

func createPlayer(gameMap [][]uint) (player player) {
	player.positionPlayer(gameMap)

	player.dirX = -1
	player.dirY = 0
	player.planeX = 0
	player.planeY = 0.66
	player.movementSpeed = 1
	player.rotationSpeed = 3

	return player
}

func (p *player) positionPlayer(gameMap [][]uint) {
	for y := range len(gameMap) {
		for x := range len(gameMap[y]) {
			if x == 0 || y == 0 {
				continue
			}

			if gameMap[y][x] == 0 {
				p.positionX = float64(x)
				p.positionY = float64(y)
				return
			}
		}
	}
}
