package main

import (
	"fmt"
	"time"
)

func out(from, to int, ch chan bool) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
	ch <- true
}

func evenSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	ch <- result
}
func squareSum(from, to int, ch chan int) {
	result := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			result += i * i
		}
	}
	ch <- result
}

func main() {
	ch := make(chan bool)

	go out(0, 5, ch)
	go out(6, 10, ch)

	<-ch
	<-ch

	evenCh := make(chan int)
	sqCh := make(chan int)

	go evenSum(0, 100, evenCh)
	go squareSum(0, 100, sqCh)
    
	// select {
	// case x := <-evenCh:
	// 	fmt.Println(x)
	// case y := <-sqCh:
	// 	fmt.Println(y)
	// default:
	// 	fmt.Println("Ничего не доступно")
	//  time.Sleep(50 * time.Millisecond)
	// }
	fmt.Println(<-evenCh + <-sqCh)
}
