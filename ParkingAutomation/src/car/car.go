package car

import "fmt"

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
