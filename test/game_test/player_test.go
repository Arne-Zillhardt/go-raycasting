package game_test

import (
	"log"
	"testing"

	"github.com/arne-zillhardt/raycasting/pkg/game"
)

func TestPlayerCreation(t *testing.T) {
	log.Println("TEST: player creation test")
	gameToTest := game.GetInstance()

	playerPositionX, playerPositionY := gameToTest.PlayerPosition()
	if playerPositionX != 1 {
		log.Println("Start-PositionX not 1")
		t.Fail()
	}
	if playerPositionY != 1 {
		log.Println("Start-PositionY not 1")
		t.Fail()
	}
}
