package main

import "fmt"

//func helloChannel(done chan bool)  {
//	fmt.Println("goroutine is going to sleep")
//	time.Sleep(4 * time.Second)
//	fmt.Println("hello world goroutine")
//	done <- true
//}

func calcSquares(num int, squareop chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	squareop <- sum
}

func calcCubes(num int, cubeop chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	cubeop <- sum
}

//unidirectional channels
func sendData(sendch chan <-int)  {
	sendch <- 10
}

//closing and looping channels
func producer(chnl chan int)  {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	//declaring channels
	//var a chan int
	//
	//if a == nil {
	//	fmt.Println("channel a is nil")
	//	a = make(chan int)
	//	fmt.Printf("type of a is %T", a)
	//}

	//done := make(chan bool)
	//fmt.Println("main going to call hello go goroutine")
	//go helloChannel(done)
	//<- done
	//fmt.Println("main function called")

	//num := 589
	//sqrch := make(chan int)
	//cubech := make(chan int)
	//go calcSquares(num, sqrch)
	//go calcCubes(num, cubech)
	//squares, cubes := <-sqrch, <-cubech
	//fmt.Println("final output", squares + cubes)

	//deadlock
	//ch := make(chan int)
	//ch <- 5

	//unidirectional channels
	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)

	//closing and looping channels
	ch := make(chan int)
	go producer(ch)
	//for  {
	//	v, ok := <-ch
	//	if ok == false {
	//		break
	//	}
	//	fmt.Println("received ", v, ok)
	//}
	for v := range ch{
		fmt.Println("received ", v)
	}
}

//sending form a channel
// data := <- a

//receiving from a channel
// a <- data
