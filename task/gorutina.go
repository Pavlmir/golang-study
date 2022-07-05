package main

import (
	"fmt"
	"time"
)

func say(m []int) {
	for _, el := range m {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(el)
	}
}

func main() {
    var a = []int{1, 2, 3, 4, 5}
	var b = []int{6, 7, 8, 9, 10}
	go say(a)
	say(b)
}