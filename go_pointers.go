package main

import "fmt"

func main()  {
	//b := 255
	//var a *int = &b
	//fmt.Printf("type of a is %T\n", a)
	//fmt.Println("address of b is", a)
	//
	//x := 25
	//var y *int
	//if y == nil {
	//	fmt.Println("y is ", y)
	//	y = &x
	//	fmt.Println("y after initialization is", y)
	//}

	//size := new(int)
	//fmt.Printf("size value %d, type is %T, address is %v\n", size, size, size)
	//*size = 25
	//fmt.Println("new size value", *size)

	//b := 255
	//a := &b
	//
	//fmt.Println("address of b is", a)
	//fmt.Println("value of b is", *a)
	//*a++
	//fmt.Println("new value of b is", b)

	//a := 58
	//fmt.Println("value of a before function call is", a)
	//b := &a
	//value(b)
	//fmt.Println("value after function call", a)

	d := pointerFromFunc()
	fmt.Println("value from function", *d)

	a := [3]int{89, 90, 91}
	modify(a[:])
	fmt.Println(a)

}

func value(val *int)  {
	*val = 55
}

func pointerFromFunc() *int {
	i := 5
	return &i
}

func modify(arr []int)  {
	arr[0] = 90
}