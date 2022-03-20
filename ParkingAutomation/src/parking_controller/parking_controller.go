package parking_controller

import "parking/src/car"

var ParkingList []string

func Check(location string) bool {
	check := false
	for _,item := range ParkingList {
		if item == location {
			return true
		}
	}
	return check
}

func Add(location string){
	ParkingList = append(ParkingList, location)
}

func Delete(location string) {
	tmpList := []string{}
	tmpCarList := []car.Car{}
	for _,item := range ParkingList {
		if item != location {
			tmpList = append(tmpList, item)
		}
	}

	for _,item := range car.CarList {
		if item.Car_location != location {
			c := car.Car{}
			c.Car_location = item.Car_location
			c.Number_plate = item.Number_plate
			tmpCarList = append(tmpCarList, c)
		}
	}
	car.CarList = tmpCarList
	ParkingList = tmpList
}
