package main

import "fmt"

func number() int {
	num := 15 * 5
	return num
}

func main()  {
	//finger := 7
	//
	//switch finger {
	//case 1:
	//	fmt.Println("Thumb")
	//case 2:
	//	fmt.Println("Index")
	//case 3:
	//	fmt.Println("Rink")
	//case 4:
	//	fmt.Println("Ring")
	//case 5:
	//	fmt.Println("Pinky")
	//default:
	//	fmt.Println("Incorrect number")
	//}

	// multiple expression

	//letter := "i"
	//
	//switch letter {
	//case "a", "e", "u", "i", "o":
	//	fmt.Println("vowel")
	//default:
	//	fmt.Println("letter not found")
	//}

	// expressionless switch

	//num := 75
	//switch {
	//case num >= 0 && num <= 50:
	//	fmt.Println("number is greater than zero and less than 50")
	//case num >= 51 && num <= 100:
	//	fmt.Println("num is greater than 51 and less than 100")
	//case num >= 101:
	//	fmt.Println("num is greater than 100")
	//}

	switch num := number(); {
	case num < 50:
		fmt.Println("less than 50")
	case num < 100:
		fmt.Println("less than 100")
		fallthrough
	case num < 200:
		fmt.Println("less than 200")
	}
}
