package loops

func main() {

	//for loop is the only loop in Go, and it has three components separated by semicolons: the init statement, the condition expression and the post statement.
	//The init statement will always execute before the first iteration.
	//The condition expression is evaluated at the beginning of every iteration.
	//The post statement will always execute at the end of every iteration.

	for i := 0; i < 5; i++ {
		println(i)
	}

	//The init and post statements are optional.
	//If you omit the init statement, it is like you have an infinite loop with no initialization and the condition expression is always true.
	//If you omit the post statement, it is like you have an infinite loop with no increment and the condition expression is always true.

	j := 0
	for j < 5 {
		println(j)
		j++
	}

	//If you omit the condition expression, it is like you have an infinite loop with no condition and the condition expression is always true.

	k := 0
	for {
		println(k)
		k++
		if k == 5 {
			break
		}
	}

	//The break statement is used to terminate the for loop.
	//The continue statement is used to skip the current iteration and move to the next iteration.

	for l := 0; l < 5; l++ {
		if l == 2 {
			continue
		}
		println(l)
	}

	//The break and continue statements can be used with any for loop.
	//Continue statement will skip the current iteration and move to the next iteration.

	for m := 0; m < 5; m++ {
		for n := 0; n < 2; n++ {
			if m == n {
				continue
			}
			println(m, n)
		}
	}

	// with range keyword example
	//The range keyword is used in for loop to iterate over items of an array, slice, channel or map.
	//With array and slices, it returns the index and value of the item.
	//With maps, it returns the key and value of the item.
	//With strings, it returns the index and byte value of the item.

	//with array
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		println(i, v)
	}

}
