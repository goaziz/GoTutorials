package main

import (
	"GoTutorials/rectangle"
	"fmt"
	"log"
)
//Packages are used to organise go source code for better reusability and readability.

var rectLen, recWidth float64 = 6, 7

func init()  {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if recWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main()  {
	fmt.Println("Go packages in this tutorial")
	fmt.Printf("area pf rectangle %.2f\n", rectangle.Area(rectLen, recWidth))
	fmt.Printf("area pf rectangle %.2f\n", rectangle.Diagonal(rectLen, recWidth))
}
