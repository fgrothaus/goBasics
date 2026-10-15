package main

import (
	"opc_ua/basics"
)

func main() {
	// temp, err := basics.LeseTemperatur()

	// if err != nil {
	// 	fmt.Println("Fehler: ", err)
	// 	return
	// }
	// fmt.Println("Temperatur:", temp)

	basics.GoRoutines()
}
