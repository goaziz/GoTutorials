package main

import (
	"fmt"
	"math"
)

func main() {
	const j = 55
	fmt.Println(j)

	fmt.Println("Hello world")
	var i = math.Sqrt(64)
	fmt.Println(i)

	const a = 7
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a

	fmt.Println("intvar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

}