package main

import (
	"fmt"
)

//func changeLocal(num [5]int)  {
//	num[0] = 55
//	fmt.Println("inside function", num)
//}

func main()  {
	//num := [...]int{5, 4, 3, 2, 1}
	//fmt.Println("before passing function", num)
	//changeLocal(num)
	//fmt.Println("after passing function", num)

	//a := [...]int{12, 54, 76}
	//fmt.Println(a)

	//a := [...]float64{56.2, 12.5, 54.7, 56}
	//fmt.Println("length of the array", len(a))
	//
	//for i := 0; i < len(a); i++ {
	//	fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	//}
	//
	//sum := float64(0)
	//for _, i := range a{
	//	sum += i
	//}
	//fmt.Println("\nsum of numbers in the list is", sum)
	//
	//s := [3][2]string{
	//	{"lion", "tiger"},
	//	{"cat", "dog"},
	//	{"pigeon", "peacock"},
	//}
	//printArray(s)
	//var b[3][2]string
	//b[0][0] = "apple"
	//b[0][1] = "samsung"
	//b[1][0] = "google"
	//b[1][1] = "u-cell"
	//b[2][0] = "beeline"
	//b[2][1] = "mobiuz"
	//fmt.Printf("\n")
	//printArray(b)

	c := [5]int{6, 7, 8, 2, 9}
	var b []int = c[1:4]
	fmt.Println(b)
}

//func printArray(s [3][2]string)  {
//	for _,v1 := range s {
//		for _, v2 := range v1 {
//			fmt.Printf("%s ",v2)
//		}
//		fmt.Printf("\n")
//	}
//}