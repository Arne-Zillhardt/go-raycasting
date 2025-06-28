package player_test

import (
	"log"
	"testing"

	"github.com/arne-zillhardt/raycasting/pkg/player"
)

func TestPlayerCreation(t *testing.T) {
	log.Println("TEST: player creation test")
	playerToTest := player.GetInstance()

	if playerToTest.GetPositionX() != 1 {
		log.Println("Start-PositionX not 1")
		t.Fail()
	}
	if playerToTest.GetPositionY() != 1 {
		log.Println("Start-PositionY not 1")
		t.Fail()
	}
}

func TestPlayerMoveForward(t *testing.T) {
	log.Println("TEST: player movement forward test")
	playerToTest := player.GetInstance()

	playerToTest.Move(player.FORWARD)

	if playerToTest.GetPositionX() != 1 {
		log.Println("After moving forward, PositionX should not change")
		t.Fail()
	}
	if playerToTest.GetPositionY() != 0 {
		log.Println("PositionY after moving forward is not 0")
		t.Fail()
	}
}
