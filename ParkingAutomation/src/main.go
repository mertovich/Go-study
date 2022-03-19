package main

import (
	"fmt"
	"parking/src/car"
	"parking/src/errors"
	"parking/src/parking_controller"
)

func main() {
	fmt.Println(`--- Start ---`)

L:
	for true {
		var choice string
		// menu
		fmt.Println(`[1] Register`)
		fmt.Println(`[2] Search`)
		fmt.Println(`[3] All Records`)
		fmt.Println(`[4] Exit`)
		fmt.Print(`Choice: `)
		fmt.Scan(&choice)

		// menu func
		switch choice {
		case `1`:
			fmt.Println(`Case 1`)
			// read value
			numberPlate := ``
			carLocation := ``
			fmt.Print(`Number plate: `)
			fmt.Scan(&numberPlate)
			fmt.Print(`Car location: `)
			fmt.Scan(&carLocation)
			// parking chech
			parkingcheck := parking_controller.Check(carLocation)
			if parkingcheck == false {
				c := car.Car{numberPlate, carLocation}
				car.Add(c)
				parking_controller.Add(carLocation)
			} else {
				errors.ParkingLotFull(carLocation)
			}
		case `2`:
			fmt.Println(`Case 2`)
			break L
		case `3`:
			fmt.Println(`Case 3`)
			car.GetAllList()
		case `4`:
			fmt.Println(`Exit`)
			break L
		default:
			errors.IncorrectEntry()
			break L
		}

	}
}
