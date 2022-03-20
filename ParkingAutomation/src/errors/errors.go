package errors

import "fmt"

func IncorrectEntry() {
	fmt.Println(`incorrect entry`)
}

func ParkingLotFull(location string) {
	fmt.Println(`Parking lot full `,location)
}

func RegisterNull() {
	fmt.Println(`Register null`)
}
