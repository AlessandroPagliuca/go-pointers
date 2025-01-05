package main

import "fmt"

func main() {
	age := 32 //regular variable

	//pointer variable
	agePointer := &age

	fmt.Println("Age: ", *agePointer)

	editAgeAdultAge(agePointer)
	fmt.Println(age)
}

func editAgeAdultAge(age *int) {
	*age = *age - 18
}
