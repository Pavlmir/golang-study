package main

import "fmt"

// В этой функции мы считываем данные из канала c и выводим в консоль.
func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!") // считываем данные из канала
}

func main() {
	fmt.Println("main() started") // В функции main программа сначала выводит "main() started".
	c := make(chan string) // это небуферизованный канал, ожидаем считывание из канала

	go greet(c) // запускаем функцию как горутину, используя ключевое слово go.

	c <- "John" // запись в канал
	fmt.Println("main() stopped") // main блокируется до тех пор, пока другая горутина (greet) не считает данные из канала c
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main1() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
