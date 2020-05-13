package main

import "fmt"

func main() {
	//make(map[string]float32)
	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["Jess"] = 32

	fmt.Println(grades)

	TimsGrade := grades["Timmy"]
	delete(grades, "Timmy")
	fmt.Println(TimsGrade)
	fmt.Println(grades)

	for key, value := range grades {
		fmt.Printf("%s:%f", key, value)
	}
}
