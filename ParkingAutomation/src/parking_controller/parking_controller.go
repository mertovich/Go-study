package parking_controller

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
