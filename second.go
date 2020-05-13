package main

import (
	"fmt"
)

/*
func foo() {
	fmt.Println("Square root", math.Sqrt(5))
}
*/

//godoc fmt Println

/*
func main() {
	fmt.Println("Random number 1-100:", rand.Intn(100))
}
*/

// x,y float64 as function parameters

func add(x float64, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	//var num1 float64 = 5.6
	//var num2 float64 = 9.5

	var num1, num2 float64 = 5.6, 9.5

	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))
	//fmt.Println(add(num1, num2))
}

//If var is not used for declaration  := should be used

//Gives error for not used variables not warning
//The type cannot be changed once defined i.e. once the program is run

//const x int = 5
