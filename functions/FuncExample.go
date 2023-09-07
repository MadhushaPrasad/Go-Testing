package functions

import "fmt"

// FuncExample what is a function?
// a function is a block of code that performs a specific task
// a function is a reusable block of code
// a function is a collection of statements that are grouped together to perform an operation
// without function call, a function will not execute

// syntax for declaring a function => func keyword -> function name -> function body
// func function_name() {
// 	// function body
// }

func funcExample() {

}

// syntax for declaring a function with parameters => func keyword -> function name -> parameters -> function body
// func function_name(parameters) {
// 	// function body
// }

func funcExampleWithParameters(firstName string, lastName string) {
	//print the full name
	fmt.Println(firstName + " " + lastName)
}

// syntax for declaring a function with return type => func keyword -> function name -> return type -> function body
// func function_name() return_type {
// 	// function body
// }

func funcExampleWithReturnType() string {
	return "Hello World"
}

// syntax for declaring a function with parameters and return type => func keyword -> function name -> parameters -> return type -> function body
// func function_name(parameters) return_type {
// 	// function body
// }

func funcExampleWithParametersAndReturnType(firstName string, lastName string) string {
	//print the full name
	return firstName + " " + lastName
}

// syntax for declaring a function with multiple parameters and return type => func keyword -> function name -> parameters -> return type -> function body
// func function_name(parameters) return_type {
// 	// function body
// }

func funcExampleWithMultipleParametersAndReturnType(firstName string, lastName string, age int) string {
	//print the full name
	return firstName + " " + lastName + " " + string(age)
}

// syntax for declaring a function with multiple parameters and multiple return type => func keyword -> function name -> parameters -> return type -> function body
// func function_name(parameters) return_type {
// 	// function body
// }

func funcExampleWithMultipleParametersAndMultipleReturnType(firstName string, lastName string, age int) (string, string, int) {
	//print the full name
	return firstName, lastName, age
}

//call all above functions for execution

func main() {
	funcExample()
	funcExampleWithParameters("John", "Doe")
	fmt.Println(funcExampleWithReturnType())
	fmt.Println(funcExampleWithParametersAndReturnType("John", "Doe"))
	fmt.Println(funcExampleWithMultipleParametersAndReturnType("John", "Doe", 30))
	fmt.Println(funcExampleWithMultipleParametersAndMultipleReturnType("John", "Doe", 30))
}
