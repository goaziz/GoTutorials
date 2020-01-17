package main

import "fmt"

func main()  {
	num := 10

	if num % 2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	x := 99
	if x <= 50 {
		fmt.Println("number is less than or equal to 50")
	}else if x >= 50 && x <= 100 {
		fmt.Println("number is between 51 and 100")
	}else {
		fmt.Println("number is greater than 100")
	}
}