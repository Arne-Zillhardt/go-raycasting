package game

import (
	"log"
)

type player struct {
	positonX  *int32
	positonY  *int32
	viewAngle *float64
}

type movementDirection struct {
	additionToPositionX int16
	additionToPositionY int16
}

var (
	FORWARD  movementDirection = movementDirection{0, -1}
	BACKWARD movementDirection = movementDirection{0, 1}
	LEFT     movementDirection = movementDirection{-1, 0}
	RIGHT    movementDirection = movementDirection{1, 0}
)

func initPlayer() player {
	log.Println("Created player")
	player := player{
		positonX:  new(int32),
		positonY:  new(int32),
		viewAngle: new(float64),
	}

	*player.positonX = int32(1)
	*player.positonY = int32(1)
	*player.viewAngle = float64(90)
	return player
}

func (p player) Move(direction movementDirection) {
	log.Println("Moving the player ", direction)
	log.Println("Before moving ", *p.positonX, " ", *p.positonY)

	*p.positonX += int32(direction.additionToPositionX)
	*p.positonY += int32(direction.additionToPositionY)

	log.Println("After moving ", *p.positonX, " ", *p.positonY)
}

func (p player) position() (int, int) {
	return int(*p.positonX), int(*p.positonY)
}
