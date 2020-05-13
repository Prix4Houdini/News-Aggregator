package main

import "fmt"

func foo() {
	defer fmt.Println("Done!")
	defer fmt.Println("Are we Done")
	fmt.Println("ABC")
}

func main() {
	foo()
}
