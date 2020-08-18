// welcome back, we will be designing a calculator program today!
package main

import "fmt"

// yes, just like in python, we will define functions for addition, subtraction, multiplication and division
// also, a function must have a specified return type.
// we will only work with integers, so the return type here is int
// we also specify that num1, and num2 are integers
func add(num1, num2 int) int{
	var result int
	result = num1 + num2
	return result
}

func subtract() {

}

func divide() {

}

func multiply() {

}

func main() {
	fmt.Println("Welcome to the calculator program!")
	fmt.Println("1. Add Two Numbers")
	fmt.Println("2. Perform Subtraction")
	fmt.Println("3. Perform Multiplication")
	fmt.Println("4. Perform Division")
	fmt.Println("Please enter your choice")
	// ah yes, variables in go
	// variables are prefixed with var, usually. though this is not the only way to do this
	// then, the type of the variable must be declared. in this case, for the choice, we will use an int
	var inp int

	// the Scanln is used in place of the input() opertor. instead of saying variable = input()
	// we use fmt.Scanln() and use the &operator to specify what we want to read the value into
	fmt.Scanln(&inp)

	// now let us take the two numbers
	fmt.Println("Enter the first number")
	var first int
	fmt.Scanln(&first)
	fmt.Println("Enter the second number")
	var second int
	fmt.Scanln(&second)

	// basic if else block. there are no brackets, and the colon in python is replaced by {brackets}
	if inp == 1 {
		// let us call the function add()
		// we can also do a kind of simultaneous declaration
		// this automatically declares and assigns a type to the answer variable
		answer := add(first, second)
		fmt.Println(answer)

	} else if inp == 2 {
		// important with the else if
		// the else must be directly inline with the closing bracket with the if
		// else you'll get an unexpected else error

	} else if inp == 3 {

	} else if inp == 4 {

	}

}