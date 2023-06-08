package main

import "fmt"

// Struct definition
type Person struct {
	name string
	age  int
}

// Function to update age by value
func updateAgeByValue(age int) {
	age = age + 1
}

// Function to update age by reference
func updateAgeByReference(age *int) {
	*age = *age + 1
}

func main() {
	// Variable declaration and initialization
	var age int = 15
	var name string = "Vishal"

	// Create a Person struct instance
	person := Person{
		name: name,
		age:  age,
	}

	// Print the struct values
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)

	// Call the function to update age by value
	updateAgeByValue(age)
	fmt.Println("Age after update by value:", age) // Value of age remains unchanged

	// Call the function to update age by reference
	updateAgeByReference(&age)
	fmt.Println("Age after update by reference:", age) // Value of age is updated

	// Print a message using the fmt package
	fmt.Println("Hello, World!")
}
