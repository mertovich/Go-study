package car

import (
	"fmt"
	"parking/src/errors"
)

type Car struct {
	Number_plate string
	Car_location string
}

// Global variable
var CarList []Car

func Add(c Car){
	CarList = append(CarList,c)
}

func GetAllList() {
	for _, car := range CarList {
		fmt.Println(`Plate: `,car.Number_plate,`Location: `,car.Car_location)
	}
}

func SearchByPlate(plate string) {
	tmpCar := Car{}
	for _,item := range CarList {
		if item.Number_plate == plate {
			tmpCar.Car_location = item.Car_location
			tmpCar.Number_plate = item.Number_plate
		}
	}
	if tmpCar.Car_location != `` && tmpCar.Number_plate != ``{
		fmt.Println(`Plate: `,tmpCar.Number_plate,`Location: `,tmpCar.Car_location)
	} else {
		errors.RegisterNull()
	}
}

func SearchByLocation(location string) {
	tmpCar := Car{}
	for _,item := range CarList {
		if item.Number_plate == location {
			tmpCar.Car_location = item.Car_location
			tmpCar.Number_plate = item.Number_plate
		}
	}
	if tmpCar.Car_location != `` &&  tmpCar.Number_plate != `` {
		fmt.Println(`Plate: `,tmpCar.Number_plate,`Location: `,tmpCar.Car_location)
	} else {
		errors.RegisterNull()
	}
}
