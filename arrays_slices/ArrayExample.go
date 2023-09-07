package arrays_slices

import "fmt"

/*
//Arrays

An array is a collection of elements that belong to the same type. For example,
the collection of integers 5, 8, 9, 79, 76 forms an array. Mixing values of
different types, for example, an array that contains both strings and integers
is not allowed in Go.

//Declaration

An array belongs to type [n]T. n denotes the number of elements in an array
and T represents the type of each element. The number of elements n is also
a part of the type(We will discuss this in more detail shortly.)

There are different ways to declare arrays. Let's look at them one by one.*/

func main() {
	var a [3]int //int array with length 3
	fmt.Println(a)

	var b [3]int //int array with length 3
	b[0] = 12    // array index starts at 0
	b[1] = 78
	b[2] = 50

	c := [3]int{12, 78, 50} // shorthand declaration to create array
	// Go has a short variable declaration operator :=
	//it declares a variable and assigns a value in one step.

	d := [...]int{12, 78, 50} // ... makes the compiler determine the length

	fmt.Println(c)
	fmt.Println(d)
}
