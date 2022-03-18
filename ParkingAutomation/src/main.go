package main

import (
	"fmt"
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
		fmt.Scan(&choice)

		// menu func
		switch choice {
		case `1`:
			fmt.Println(`Case 1`)
			break L
		case `2`:
			fmt.Println(`Case 2`)
			break L
		case `3`:
			fmt.Println(`Case 3`)
			break L
		default:
			break L
		}

	}
}
