package main

import "fmt"

type car struct {
	gas_pedal      uint16 //0 to 65535
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func main() {
	a_car := car{145, 0, 142, 255} //You can specifically name(but all should if done)

	fmt.Println(a_car.gas_pedal)
}
