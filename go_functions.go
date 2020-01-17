package main

import "fmt"

//A function is a block of code that performs a specific task
//A function takes a input, performs
//some calculations on the input and generates a output.

//The syntax of the function

// func functionname(parameter type) returntype {
//		function body
//}

//func calculateBill(price, no int) int {
//	var totalPrice = price * no
//	return totalPrice
//}
//
//func main()  {
//	price, no := 90, 6
//	totalPrice := calculateBill(price, no)
//
//	fmt.Println("Total price", totalPrice)
//}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main()  {
	//area, perimeter := rectProps(10.8, 5.3)
	//fmt.Printf("Area %f, Perimeter %f", area, perimeter)

	//blank identifier known as '_' can be used in place of any value of any type

	area, _ := rectProps(10.8, 3.1)
	fmt.Printf("the are is %f", area)
}