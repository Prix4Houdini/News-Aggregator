package main

import "fmt"

func main() {

	i := 0
	for i < 10 {
		fmt.Println(i)
		i += 5
	}

	//Infinite Loop

	for {
		fmt.Println("STUFFS")
	}

	x := 5
	for {
		fmt.Println("Stiff", x)
		x += 3
		if x > 25 {
			break
		}
	}
	a := 3
	for x := 5; a < 25; x += 3 {
		fmt.Println("Stuff")
		a += 3
	}
}
