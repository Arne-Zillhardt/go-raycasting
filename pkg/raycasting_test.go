package pkg

import (
	"log"
	"math"
	"testing"
)

func TestSimpleRayCalculation(t *testing.T) {
	log.Println("Test the calculation of a ray")

	dirX, dirY := -1.0, 0.0
	planeX, planeY := 0.0, 0.66

	testMap := getTestMap()
	player := player{}
	player.positionX = 1.0
	player.positionY = 2.0

	w := 640
	x := 320

	cameraX := 2 * float64(x) / float64(w) - 1
	var rayDirX float64 = dirX + planeX * cameraX
	var rayDirY float64 = dirY + planeY * cameraX

	distance, side := calculateRayLength(player, rayDirX, rayDirY, testMap)

	if distance != 1 {
		log.Println("Distance: ", distance)
		log.Println("Expected the distance to be 1")
		t.Fail()
	}

	if side != 0 {
		log.Println("Side: ", side)
		log.Println("Expected the side to be 0")
		t.Fail()
	}

	dirX = 1

	rayDirX = dirX + planeX * cameraX
	rayDirY = dirY + planeY * cameraX

	distance, side = calculateRayLength(player, rayDirX, rayDirY, testMap)

	if distance != 22 {
		log.Println("Distance: ", distance)
		log.Println("Expected the distance to be 22")
		t.Fail()
	}

	if side != 0 {
		log.Println("Side: ", side)
		log.Println("Expected the side to be 0")
		t.Fail()
	}
}

func TestCalculationOfDeltaDistance(t *testing.T) {
	log.Println("Test the calculation of the delta-distance")

	var rayDir float64 = 10

	deltaDistance := calculateDeltaDistance(rayDir)

	if deltaDistance != 0.1 {
		log.Println("Expected the delta-distance to be 0.1")
		t.Fail()
	}

	rayDir = 0

	deltaDistance = calculateDeltaDistance(rayDir)

	if deltaDistance != math.MaxFloat64 {
		log.Println("Expected the delta-distance to be the maximum value of a float")
		t.Fail()
	}
}

func TestCalculationOfSideDistanceAndStep(t *testing.T) {
	log.Println("Test the calculation of side-distance and step")

	var rayDir float64 = 1
	var playerPosition float64 = 10.5
	var deltaDistance float64 = 1

	step, sideDist := calculateSideDistanceAndStep(rayDir, playerPosition, deltaDistance)

	if step != 1 {
		log.Println("Expected the step to be 1")
		t.Fail()
	}
	if sideDist != 0.5 {
		log.Println("Expected the side-distance to be 0.5")
		t.Fail()
	}

	rayDir = -1

	step, sideDist = calculateSideDistanceAndStep(rayDir, playerPosition, deltaDistance)

	if step != -1 {
		log.Println("Expected the step to be 1")
		t.Fail()
	}
	if sideDist != 0.5 {
		log.Println("Expected the side-distance to be 0.5")
		t.Fail()
	}
}
