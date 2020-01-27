package main

import "fmt"

// Go Loops. for is only available loop in Go. Go does not have while or do wile loops.

// for loop syntax
// for initialisation; condition; post {
//}

func main()  {
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i)
	//}

	// break
	//for i := 1; i <= 10; i++ {
	//	if i > 5 {
	//		break
	//	}
	//	fmt.Println(i)
	//}

	// continue
	//for i := 1; i <= 10; i++ {
	//	if i % 2 == 0{
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	
	// nested loops
	//n := 5
	//for i := 0; i < n; i++ {
	//	for j := 0; j <= i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}
	
	// labels
	outer:
		for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++ {
				fmt.Printf("i = %d, j = %d\n", i, j)

				if i == j {
					break outer
				}
			}
		}

	i := 0
	for i <= 10 {
		fmt.Printf("%d\n", i)
		i += 2
	}
}
