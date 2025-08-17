package pkg

import (
	"log"
	"testing"
)

func TestCreationOfMap(t *testing.T) {
	width, height := uint(24), uint(24)
	instanceToTest := getGameMap(width, height)

	if len(instanceToTest) != int(width) || len(instanceToTest[0]) != int(height) {
		log.Println("Expected the map to be width x height")
		t.Fail()
	}

	for y := range(height) {
		for x := range(width) {
			if instanceToTest[x] != 1 || 
		}
	}
}
