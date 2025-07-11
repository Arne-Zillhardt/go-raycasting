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

	instanceToTest.changeView(INCREASE, 0.1)

	if *instanceToTest.viewAngle != 90.5 {
		log.Println("Incorrect initial viewAngle. Expected 90.5 but was ", *instanceToTest.viewAngle)
		t.Fail()
	}

	instanceToTest.changeView(DECREASE, 0.1)

	if *instanceToTest.viewAngle != 90 {
		log.Println("Incorrect initial viewAngle. Expected 90 but was ", *instanceToTest.viewAngle)
		t.Fail()
	}


	*instanceToTest.viewAngle = 2
	instanceToTest.changeView(DECREASE, 1)

	if *instanceToTest.viewAngle != 357 {
		log.Println("Incorrect initial viewAngle. Expected 357 but was ", *instanceToTest.viewAngle)
		t.Fail()
	}

	*instanceToTest.viewAngle = 357
	instanceToTest.changeView(INCREASE, 1)

	if *instanceToTest.viewAngle != 2 {
		log.Println("Incorrect initial viewAngle. Expected 2 but was ", *instanceToTest.viewAngle)
		t.Fail()
	}
}

func TestCalculateViewLine(t *testing.T) {
	log.Println("TEST: Test the vector calculation of viewline")

	instanceToTest := initPlayer()

	linePositionX, linePositionY := instanceToTest.calculateLinePosition(0)

	if (linePositionX < 6 || linePositionX > 6.5) || (linePositionY < 1 || linePositionY > 1.5) {
		log.Println("Line not as expected. Expected 6 and 1, got ", linePositionX, " ", linePositionY)
		t.Fail()
	}

	linePositionX, linePositionY = instanceToTest.calculateLinePosition(90)

	if (linePositionX < 1 || linePositionX > 1.5) || (linePositionY > -3.5 || linePositionY < -4.5) {
		log.Println("Line not as expected. Expected 1 and -4, got ", linePositionX, " ", linePositionY)
		t.Fail()
	}

	linePositionX, linePositionY = instanceToTest.calculateLinePosition(-180)

	if (linePositionX > -3.5 || linePositionX < -4.5) || (linePositionY < 1 || linePositionY > 1.5) {
		log.Println("Line not as expected. Expected -4 and 1, got ", linePositionX, " ", linePositionY)
		t.Fail()
	}
}

func TestMove(t *testing.T) {
	log.Println("TEST: Test player movement")

	instanceToTest := initPlayer()

	if *instanceToTest.positonX != 1 || *instanceToTest.positonY != 1 {
		log.Println("Incorrect start position")
		t.Fail()
	}

	instanceToTest.move(RIGHT, 0.1)

	if *instanceToTest.positonX != 1.1 || *instanceToTest.positonY != 1 {
		log.Println("Expected 1.1 on x, but got ", *instanceToTest.positonX, " ", *instanceToTest.positonY)
		t.Fail()
	}
	
	instanceToTest.move(LEFT, 0.1)

	if *instanceToTest.positonX != 1 || *instanceToTest.positonY != 1 {
		log.Println("Expected 1 on x, but got ", *instanceToTest.positonX, " ", *instanceToTest.positonY)
		t.Fail()
	}

	instanceToTest.move(BACKWARD, 0.1)

	if *instanceToTest.positonX != 1 || *instanceToTest.positonY != 1.1 {
		log.Println("Expected 1.1 on y, but got ", *instanceToTest.positonX, " ", *instanceToTest.positonY)
		t.Fail()
	}

	instanceToTest.move(FORWARD, 0.1)

	if *instanceToTest.positonX != 1 || *instanceToTest.positonY != 1 {
		log.Println("Expected 1 on y, but got ", *instanceToTest.positonX, " ", *instanceToTest.positonY)
		t.Fail()
	}
}
