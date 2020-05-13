package main

import "fmt"

const usixteenmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16 //0 to 65535
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

//Value Receiver method
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenmax)
}

func (c *car) mph() float64 {
	c.top_speed_kmh = 500
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenmax / kmh_multiple)
}

//Here value is permanently changed since its pointer reciever
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main() {
	a_car := car{145, 0, 142, 255} //You can specifically name(but all should if done)

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}
