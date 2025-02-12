package main

import (
	"fmt"

	"github.com/lrbell17/gotutorial/sec-09-interfaces/stub/exercise02/acura"
	"github.com/lrbell17/gotutorial/sec-09-interfaces/stub/exercise02/ford"
	"github.com/lrbell17/gotutorial/sec-09-interfaces/stub/exercise02/honda"
	"github.com/lrbell17/gotutorial/sec-09-interfaces/stub/exercise02/tesla"
	"github.com/lrbell17/gotutorial/sec-09-interfaces/stub/exercise02/vehicle"
)

func main() {
	myGarage := make([]vehicle.Vehicle, 4)
	myGarageErr := make([]error, 4)

	myGarage[0], myGarageErr[0] = ford.NewExplorer(1995)
	myGarage[1], myGarageErr[1] = honda.NewPilot(2010)
	myGarage[2], myGarageErr[2] = acura.NewTSX(2006, vehicle.Hybrid)
	myGarage[3], myGarageErr[3] = tesla.NewModel3(2018)

	for _, e := range myGarageErr {
		if e != nil {
			fmt.Println(e)
		}
	}

	for _, v := range myGarage {
		fmt.Println(v)
	}
}
