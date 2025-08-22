package pkg

import (
	"log"
	"testing"
)

func TestCreationOfMap(t *testing.T) {
	log.Println("TEST: Test the creation of the gameMap")
	width, height := uint(24), uint(24)
	instanceToTest := getGameMap(width, height)

	if len(instanceToTest) != int(width) || len(instanceToTest[0]) != int(height) {
		log.Println("Expected the map to be ", width, "x", height)
		t.Fail()
	}

	for y := range(height) {
		if instanceToTest[y][0] != 1 || instanceToTest[y][width - 1] != 1 {
			log.Println("Expected a wall verticly")
			t.Fail()
			break
		}
	}

	for x := range(width) {
		if instanceToTest[0][x] != 1 || instanceToTest[height - 1][x] != 1 {
			log.Println("Expected a wall horizontally")
			t.Fail()
			break
		}
	}

	sideWalls := int((width * 2) + (height * 2) - 4)
	allWalls := 0
	for y := range(height) {
		for x := range(width) {
			if instanceToTest[y][x] == 1 {
				allWalls++
			}
		}
	}
	if allWalls == sideWalls {
		log.Println("Expected more walls")
		t.Fail()
	}
}
