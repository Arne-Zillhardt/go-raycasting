package pkg

import (
	"log"
	"testing"
)

func TestPlayerCreation(t *testing.T) {
	log.Println("TEST: Test the creation of player")
	instanceToTest := createPlayer(getTestMap())

	invalidX := instanceToTest.positionX != 1.0
	invalidY := instanceToTest.positionY != 2.0
	if invalidX || invalidY {
		log.Println("Invalid location")
		log.Println("PositionX: ", instanceToTest.positionX)
		log.Println("PositionY: ", instanceToTest.positionY)
		t.Fail()
	}

	if instanceToTest.dirX != -1 || instanceToTest.dirY != 0 {
		log.Println("Expected other rotation")
		log.Println("DirX: ", instanceToTest.dirX)
		log.Println("DirY: ", instanceToTest.dirY)
		t.Fail()
	}


	if instanceToTest.planeX != 0 || instanceToTest.planeY != 0.66 {
		log.Println("Expected other plane")
		log.Println("PlaneX: ", instanceToTest.planeX)
		log.Println("PlaneY: ", instanceToTest.planeY)
		t.Fail()
	}
}

func TestPlayerMovement(t *testing.T) {
	log.Println("TEST: Test the movement of player")
	instanceToTest := createPlayer(getTestMap())
	instanceToTest.positionX = 1
	instanceToTest.positionY = 2
	deltaTime := 1.0

	instanceToTest.move(FORWARD, deltaTime)

	if instanceToTest.positionX != 1 || instanceToTest.positionY != 1 {
		log.Println("Expected other position")
		log.Println("PositionX: ", instanceToTest.positionX)
		log.Println("PositionY: ", instanceToTest.positionY)
		t.Fail()
	}
}

func TestPlayerRotation(t *testing.T) {
	log.Println("TEST: Test the movement of player")
	instanceToTest := createPlayer(getTestMap())
	deltaTime := 5.0
	instanceToTest.dirX = -1
	instanceToTest.dirY = 0

	instanceToTest.rotateRight(deltaTime) Apprenticeship
	
	if instanceToTest.dirX != 1 || instanceToTest.dirY != 0 {
		log.Println("Expected other rotation")
		log.Println("DirX: ", instanceToTest.dirX)
		log.Println("DirY: ", instanceToTest.dirY)
		t.Fail()
	}
	if instanceToTest.planeX != 0 || instanceToTest.planeY != 0.66 {
		log.Println("Expected other rotation")
		log.Println("PlaneX: ", instanceToTest.planeX)
		log.Println("PlaneY: ", instanceToTest.planeY)
		t.Fail()
	}
}

func getTestMap() [][]uint {
	var gameMap [][]uint = [][]uint{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 2, 2, 2, 2, 0, 0, 0, 0, 3, 0, 3, 0, 3, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 3, 0, 0, 0, 3, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 2, 2, 0, 2, 2, 0, 0, 0, 0, 3, 0, 3, 0, 3, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 0, 0, 0, 0, 5, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 0, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 4, 4, 4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	return gameMap
}
