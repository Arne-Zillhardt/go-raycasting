package raycasting_test

import (
	"log"
	"testing"

	"github.com/arne-zillhardt/raycasting/pkg/raycasting"
)

func TestRayCalculation(t *testing.T) {
	gameMap := [][]uint8 {
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}

	calculatedLength := raycasting.CalculateRay(1, 1, 1, 4, gameMap)
	if calculatedLength != 6 {
		log.Println("Expected a length of 6, got ", calculatedLength)
		t.Fail()
	}
}
