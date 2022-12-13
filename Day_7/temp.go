package main

import "fmt"

func temp() {
	simple_int := 5
	pointer_int := &simple_int

	fmt.Println("Value of simple_int: ", simple_int)
	fmt.Println("Value of pointer_int: ", pointer_int)
	// fmt.Println()
	fmt.Printf("Address of simple_int: %p\n", &simple_int)
	fmt.Printf("Address of pointer_int: %p\n", &pointer_int)
	fmt.Println()

	changeValue(pointer_int)
	fmt.Println("After change of value")
	fmt.Println("Value of simple_int: ", simple_int)
	fmt.Println("Value of pointer_int: ", *pointer_int)

}

func changeValue(pointer *int) {
	fmt.Println("In change value")
	fmt.Printf("Address of pointer: %p\n", &pointer)
	fmt.Println("Value of pointer: ", pointer)
	fmt.Println("Actual Value of pointer: ", *pointer)

	another_value := 15
	pointer = &another_value
	fmt.Println()
	fmt.Printf("Address of pointer: %p\n", &pointer)
	fmt.Println("Value of pointer: ", pointer)
	fmt.Println("Actual Value of pointer: ", *pointer)
	fmt.Println()
}
