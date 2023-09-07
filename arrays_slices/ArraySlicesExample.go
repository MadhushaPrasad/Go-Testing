package arrays_slices

import "fmt"

/*
//Slices

A slice is a convenient, flexible and powerful wrapper on top of an array.
Slices do not own any data on their own. They are just references to existing arrays.

//Creating a slice

A slice with elements of type T is represented by []T
*/

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)

	c := []int{6, 7, 8} //creates and array and returns a slice reference

	print(c)

	/*

		//modifying a slice

		A slice does not own any data of its own. It is just a representation of the underlying array.
		Any modifications done to the slice will be reflected in the underlying array.

	*/

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dSlice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dSlice {
		dSlice[i]++
	}
	fmt.Println("array after", darr)

	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	//creating a slice using make
	i := make([]int, 5, 5)
	fmt.Println(i)

	//Appending to a slice

	/*
		As we already know arrays are restricted to fixed length and their length cannot be increased.
		Slices are dynamic and new elements can be appended to the slice using append function.
		The definition of append function is func append(s []T, x ...T) []T.
	*/

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	//Multidimensional slices
	//Similar to arrays, slices can have multiple dimensions.

	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
