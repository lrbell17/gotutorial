package ford

import (
	"errors"
	"fmt"
	"time"

	"github.com/lrbell17/gotutorial/sec-09-interfaces/stub/exercise02/vehicle"
)

type (
	// Explorer is one of Ford's vehicle that implements vehicle.Vehicle
	Explorer struct {
		vehicleData
	}
	// vehicleData is private to package
	vehicleData struct {
		year  uint16
		seats uint8
	}
)

// NewExplorer returns an initialized ford.Explorer value
func NewExplorer(year uint16) (e *Explorer, err error) {
	// verify validity of year parameter
	currentYear := uint16(time.Now().Year())
	if year < 1990 || year > currentYear {
		return nil, errors.New("invalid year for Ford Explorer")
	}

	e = &Explorer{}
	e.year = year
	e.seats = 5
	return
}

func (e *Explorer) Model() string {
	return "Explorer"
}
func (e *Explorer) Make() string {
	return "Ford"
}
func (e *Explorer) Year() uint16 {
	if e == nil {
		return 0
	}
	return e.year
}
func (e *Explorer) SeatingCap() uint8 {
	if e == nil {
		return 5
	}
	return e.seats
}
func (e *Explorer) Type() int {
	return vehicle.FOUR_DOOR_SUV
}

// TODO 1 - complete implementation of interface vehicle.Vehice for Explorer
func (e *Explorer) PowerTrain() string {
	return vehicle.Diesel.String()
}

// TODO 2 - implement fmt.Stringer for Explorer
func (e *Explorer) String() string {
	return fmt.Sprintf(
		"Make: %v, Model: %v, Year: %v, Seats: %v, Type: %v, Power Train: %v",
		e.Make(), e.Model(), e.Year(), e.SeatingCap(), vehicle.Type(e.Type()), e.PowerTrain())
}
