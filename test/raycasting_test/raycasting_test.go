package raycasting_test

import (
	"log"
	"testing"

	"github.com/arne-zillhardt/raycasting/pkg/raycasting"
)

func TestRayCalculation(t *testing.T) {
	log.Println("TEST: Test the calculation of the ray")
	gameMap := [][]uint8 {
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	if raycasting.CalculateRay(1, 1, 1, 6, gameMap) != 1 {
		log.Println("Expected a length of 1")
		t.Fail()
	}
}
