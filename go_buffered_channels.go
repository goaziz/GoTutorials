package main

import (
	"fmt"
	"sync"
	"time"
)

func write(chn chan int) {
	for i := 0; i < 5; i++ {
		chn <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(chn)
}

//wait group
func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	ch := make(chan string, 2)
	ch <- "Azizbek"
	ch <- "Abdulloh"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	/*
		chn := make(chan int, 2)
		go write(chn)
		time.Sleep(2 * time.Second)
		for v := range chn{
			fmt.Println("read value", v, "from ch")
			time.Sleep(2 * time.Second)
		}*/

	//wait group
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("all goroutines finished successfully")
}
