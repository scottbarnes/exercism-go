package elon

import "fmt"

// Car implements a remote controlled car.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// Track implements a race track.
type Track struct {
	distance int
}

// CreateCar creates a new car with given specifications.
func CreateCar(speed, batteryDrain int) *Car {
	c := &Car{batteryDrain: batteryDrain, speed: speed, battery: 100}
	return c
}

// CreateTrack creates a new track with given distance.
func CreateTrack(distance int) Track {
	t := Track{distance: distance}
	return t
}

// Drive drives the car one time.
func (car *Car) Drive() {
	if car.battery-car.batteryDrain >= 0 {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// CanFinish checks if a car is able to finish a certain track.
func (car *Car) CanFinish(track Track) bool {
	if (track.distance/car.speed)*car.batteryDrain < car.battery {
		return true
	} else {
		return false
	}
}

// DisplayDistance displays the distance the car is driven.
func (car *Car) DisplayDistance() string {
	r := fmt.Sprintf("Driven %d meters", car.distance)
	return r
}

// DisplayBattery displays the battery level.
func (car *Car) DisplayBattery() string {
	r := fmt.Sprintf("Battery at %d%%", car.battery)
	return r
}
