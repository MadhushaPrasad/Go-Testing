package pointers

import "fmt"

// syntax of a pointer is => *type
// syntax of a pointer variable is => var variable_name *type
// syntax of a pointer variable assignment is => variable_name = &value
// syntax of a pointer variable dereferencing is => *variable_name
// syntax of a pointer variable dereferencing is => *variable_name = new_value

// pointers are used to store the address of a variable and & is used to get the address of a variable
// if a pointer is not initialized, it will have a nil value
// a pointer can be initialized using the new keyword or by referencing an existing variable ex: var variable_name *type = &value

//dereferencing a pointer is used to get the value of the variable that the pointer is pointing to and * is used to dereference a pointer

func pointerExample() {
	// declare a variable
	var x int = 10

	// declare a pointer variable
	var y *int

	// assign the address of x to y
	y = &x

	// print the value of x
	fmt.Println(x)

	// print the address of x
	fmt.Println(&x)

	// print the value of y
	fmt.Println(y)

	// print the address of y
	fmt.Println(&y)

	// print the value of x using y
	fmt.Println(*y)

	// change the value of x using y
	*y = 20

	// print the value of x
	fmt.Println(x)

	// change the value of x
	x = 30

	// print the value of x using y
	fmt.Println(*y)

	// change the value of x using y
	*y = 40

	// print the value of x
	fmt.Println(x) // prints 40
}
