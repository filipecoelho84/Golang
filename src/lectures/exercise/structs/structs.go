//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

//* Create a rectangle structure containing its coordinates
type Rectangle struct {
	a Coordinate // top left
	b Coordinate // bottom right
}

func width(rect Rectangle) int {
	return (rect.b.x - rect.a.x)
}

func length(rect Rectangle) int {
	return (rect.a.y - rect.b.y)
}

//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
func area(rect Rectangle) int {
	return length(rect) * width(rect)
}

func perimeter(rect Rectangle) int {
	return (width(rect) * 2) + (length(rect) * 2)
}

//  - Print the new results to the terminal

func printInfo(rect Rectangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}

func main() {
	rect := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	//  - Print the new results to the terminal
	printInfo(rect)

	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	rect.a.y *= 2
	rect.b.x *= 2

	//  - Print the new results to the terminal
	printInfo(rect)
}

// package main

// import "fmt"

// //* Create a rectangle structure containing its coordinates
// type rectangle struct {
// 	length int
// 	width  int
// }

// //* Using functions, calculate the area and perimeter of a rectangle,
// //  - Print the results to the terminal
// //  - The functions must use the rectangle structure as the function parameter
// func area(x rectangle) {
// 	y := x.length * x.width
// 	fmt.Println(y)
// }

// func perimeter(x rectangle) {
// 	y := 2 * (x.length + x.width)
// 	fmt.Println(y)
// }

// func main() {

// 	area(rectangle{7,10})
// 	perimeter(rectangle{6,8})

// 	//* After performing the above requirements, double the size
// 	//  of the existing rectangle and repeat the calculations
// 	//  - Print the new results to the terminal

// 	area(rectangle{14,20})
// 	perimeter(rectangle{12,16})

// }
