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
		fmt.Println(`[2] Searc By Plate`)
		fmt.Println(`[3 Search By Location`)
		fmt.Println(`[4] All Parking List`)
		fmt.Println(`[5] Remove Parking `)
		fmt.Println(`[6] Exit`)
		fmt.Scan(&choice)
		// menu func
		switch choice {
		case `1`:
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
			// Plate to search
			carPlate := ``
			fmt.Print(`Plate to search :`)
			fmt.Scan(&carPlate)
			car.SearchByPlate(carPlate)
		case `3`:
			// Location to search
			carLocation := ``
			fmt.Print(`Location to search :`)
			fmt.Scan(&carLocation)
			car.SearchByLocation(carLocation)
		case `4`:
			// Get all car list
			car.GetAllList()
		case `5`:
			// Vacated location
			location := ``
			fmt.Print(`vacated location: `)
			fmt.Scan(&location)
			parking_controller.Delete(location)
		case `6`:
			// exit
			fmt.Println(`Exit`)
			break L
		default:
			errors.IncorrectEntry()
			break L
		}

	}
}
