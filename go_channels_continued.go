package main

import "fmt"

func digits(number int, dchanl chan int)  {
	for number != 0 {
		digit := number % 10
		dchanl <- digit
		number /= 10
	}
	close(dchanl)
}

func calcSquare(number int, squareop chan int)  {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch{
		sum += digit * digit
	}
	squareop <- sum
}

func calCubes(number int, cubeop chan int)  {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch{
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main()  {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquare(number, sqrch)
	go calCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("output ", squares + cubes)
}