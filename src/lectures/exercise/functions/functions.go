//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func personName(name string) string {
	fmt.Println("Hello", name)
	return name

}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func anyMessage(message string) string {
	return message
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func numbers(a, b, c int) int {
	return a + b + c
}

//* Write a function that returns any number
func anyNumber(x int) int {
	return x
}

//* Write a function that returns any two numbers
func anyTwoNumbers(y, z int) (int, int) {
	return y, z
}

func main() {

	personName("Filipe Coelho e Sara Garcia")

	fmt.Println(anyMessage("Any message!"))

	addiction := numbers(2, 3, 1)
	fmt.Println(addiction)

	randomNumber := anyNumber(37)
	fmt.Println("The chosen number was", randomNumber)

	firstRandomNumber, secondrandomNumber := anyTwoNumbers(37, 40)
	fmt.Println("The first random number is", firstRandomNumber, "\nThe second random number is", secondrandomNumber)

	//* Add three numbers together using any combination of the existing functions.
	add3Numbers := anyNumber(6) + firstRandomNumber + secondrandomNumber
	fmt.Println("The add of the 3 numbers are", add3Numbers)

}
