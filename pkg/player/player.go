package player

import (
	"log"
)

type Player interface {
	Move(movementDirection)
	GetPositionX() int32
	GetPositionY() int32
}
type player struct {
	positonX int32
	positonY int32
}

type movementDirection struct {
	additionToPositionX int16
	additionToPositionY int16
}

var (
	FORWARD movementDirection = movementDirection{0, -1}
)

func GetInstance() Player {
	log.Println("created player")
	return player{
		positonX: int32(1),
		positonY: int32(1),
	}
}

func (p player) Move(direction movementDirection) {
	log.Println("Moving the player ", direction)
	log.Println("Before moving ", p.positonX, " ", p.positonY)

	p.positonX += int32(direction.additionToPositionX)
	p.positonY += int32(direction.additionToPositionY)

	log.Println("After moving ", p.positonX, " ", p.positonY)
}

func (p player) GetPositionX() int32 {
	return p.positonX

}

func (p player) GetPositionY() int32 {
	return p.positonY
}
