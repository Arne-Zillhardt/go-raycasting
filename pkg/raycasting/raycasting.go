package raycasting

import (
	"log"
	"math"
)

func CalculateRay(x0 float64, y0 float64, x1 float64, y1 float64, field [][]uint8) float64 {
	dX := x1 - x0
	dY := y1 - y0

	steps := int(math.Abs(dY))
	if int(math.Abs(dX)) > steps {
		steps = int(math.Abs(dX))
	}

	xInc := dX / float64(steps)
	yInc := dY / float64(steps)

	x := float64(x0)
	y := float64(y0)

	hit := false
	distance := float64(1)
	for !hit {
		if int(x) >= len(field) || x < 0 || int(y) >= len(field[0]) || y < 0{
			break
		}
		if field[int(y)][int(x)] > 0 {
			hit = true
			break
		}

		x += xInc
		y += yInc
		distance++
	}

	log.Println("Distance to wall: ", distance)

	return distance
}
