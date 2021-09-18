package speed

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	c := Car{batteryDrain: batteryDrain, speed: speed, battery: 100}
	return c
}

// Track implements a race track.
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	t := Track{distance: distance}
	return t
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move but use the leftover battery.
func Drive(car Car) Car {
	switch {
	// if car will exceed battery, do nothing
	case car.batteryDrain > car.battery:
		return car
	// increment the distance and decrement the battery
	default:
		car.battery -= car.batteryDrain
		car.distance += car.speed
		return car
	}
}

// CanFinish checks if a car is able to finish a certain track.
// Could have done this with a for loop too.
func CanFinish(car Car, track Track) bool {
	if (track.distance/car.speed)*car.batteryDrain < car.battery {
		return true
	} else {
		return false
	}
}
