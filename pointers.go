package main

import "fmt"

func main() {
	age := 32 //regular variable

	fmt.Println("Age: ", age)

	adultYears := getAdultAge(age)
	fmt.Println(adultYears)
}

func getAdultAge(age int) int {
	return age - 18
}
