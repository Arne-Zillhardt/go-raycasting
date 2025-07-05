package game

import (
	"log"
	"testing"
)

func TestChangingViewingAngle(t *testing.T) {
	log.Println("TEST: Test the changing of the viewingAngle")

	instanceToTest := initPlayer()

	if *instanceToTest.viewAngle != 90 {
		log.Println("Incorrect initial viewAngle")
		t.Fail()
	}

	instanceToTest.changeView(INCREASE)

	if *instanceToTest.viewAngle != 95 {
		log.Println("Incorrect initial viewAngle. Expected 95 but was ", *instanceToTest.viewAngle)
		t.Fail()
	}

	instanceToTest.changeView(DECREASE)

	if *instanceToTest.viewAngle != 90 {
		log.Println("Incorrect initial viewAngle. Expected 90 but was ", *instanceToTest.viewAngle)
		t.Fail()
	}


	*instanceToTest.viewAngle = 2
	instanceToTest.changeView(DECREASE)

	if *instanceToTest.viewAngle != 357 {
		log.Println("Incorrect initial viewAngle. Expected 357 but was ", *instanceToTest.viewAngle)
		t.Fail()
	}

	*instanceToTest.viewAngle = 357
	instanceToTest.changeView(INCREASE)

	if *instanceToTest.viewAngle != 2 {
		log.Println("Incorrect initial viewAngle. Expected 2 but was ", *instanceToTest.viewAngle)
		t.Fail()
	}
}

func TestCalculateViewLine(t *testing.T) {
	log.Println("TEST: Test the calculation of viewline")

	instanceToTest := initPlayer()

	linePositionX, linePositionY := instanceToTest.calculateLinePosition()

	if (linePositionX < 6 || linePositionX > 6.5) || (linePositionY < 1 || linePositionY > 1.5) {
		log.Println("Line not as expected. Expected 6 and 1, got ", linePositionX, " ", linePositionY)
		t.Fail()
	}
}
