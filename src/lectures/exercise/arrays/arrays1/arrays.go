//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

//  - Products must include the price and the name
type Product struct {
	name  string
	price int
}

// func sum(products [4]Product) {
// 	for i := 0; i <= len(products); i++ {
// 		total := products[i]
// 		fmt.Println("The total cost is", total)
// 	}
// }

func main() {

	//* Using an array, create a shopping list with enough room
	//  for 4 products
	products := [...]Product{
		//* Insert 3 products into the array
		{name: "Rice", price: 2.99},
		{name: "Beans", price: 1.99},
		{name: "Meat", price: 5.99},
	}

	//* Print to the terminal:
	//  - The last item on the list
	fmt.Println("The last item on the list is", products[3-1])
	//  - The total number of items
	fmt.Println("The total number os items are", len(products))
	// //  - The total cost of the items
	// // sum(pro)
	//* Add a fourth product to the list and print out the
	//  information again
	products [4]= Product{name: "Chicken", price: 3.99}
	// // fmt.Println("My shopping list:", products)
	// // fmt.Println("The last item on the list is", products[3])

}
