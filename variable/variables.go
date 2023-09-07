package main

import "fmt"

func main() {

	/*
		var declares 1 or more variables.
		You can declare multiple variables at once.
		Go will infer the type of initialized variables.
		Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
		The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
	*/

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// syntax for declaring variables
	//var x1 int // var keyword -> variable name -> data type

	//data types examples
	var x int = 10              //int -> 32 or 64 bit  => use for assigning integer values
	var y float64 = 10.5        //float64 -> 64 bit => use for assigning decimal values
	var z string = "Hello"      //string -> 16 bit => use for assigning string values
	var a1 bool = true          //bool -> 1 bit => use for assigning boolean values
	var b1 byte = 'A'           //byte -> 8 bit => use for assigning single character values
	var c1 rune = 'B'           //rune -> 32 bit => use for assigning single character values especially for unicode characters
	var d1 complex128 = 10 + 5i //complex128 -> 128 bit => use for assigning complex values
	fmt.Println(x, y, z, a1, b1, c1, d1)
}
