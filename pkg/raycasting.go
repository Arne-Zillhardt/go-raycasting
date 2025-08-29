package pkg

import (
	"math"
)

func calculateRayLength(player player, rayDirX float64, rayDirY float64, gameMap [][]uint) (float64, int) {
	deltaDistX := calculateDeltaDistance(rayDirX)
	deltaDistY := calculateDeltaDistance(rayDirY)

	stepX, sideDistX := calculateSideDistanceAndStep(rayDirX, player.positionX, deltaDistX)
	stepY, sideDistY := calculateSideDistanceAndStep(rayDirY, player.positionY, deltaDistY)

	side := 0
	mapX := int(math.Floor(player.positionX))
	mapY := int(math.Floor(player.positionY))

	hit := 0
	for hit == 0 {
		if sideDistX < sideDistY {
			sideDistX += deltaDistX
			mapX += stepX
			side = 0
		} else {
			sideDistY += deltaDistY
			mapY += stepY
			side = 1
		}

		if gameMap[mapY][mapX] != 0 {
			hit = 1
		}
	}

	perpWallDist := sideDistX - deltaDistX
	if side == 1 {
		perpWallDist = sideDistY - deltaDistY
	}
	if perpWallDist == 0 {
		perpWallDist = 1
	}

	return perpWallDist, side
}

func calculateDeltaDistance(rayDir float64) float64 {
	deltaDist := math.Abs(1 / rayDir)
	if rayDir == 0 {
		deltaDist = math.MaxFloat64
	}

	return deltaDist
}

func calculateSideDistanceAndStep(rayDir float64, playerPosition float64, deltaDist float64) (int, float64) {
	mapPosition := math.Floor(playerPosition)

	step := 1
	sideDist := (mapPosition + 1.0 - playerPosition) * deltaDist

	if rayDir < 0 {
		step = -1
		sideDist = (playerPosition - mapPosition) * deltaDist
	}

	return step, sideDist
}
