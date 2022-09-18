package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    in := bufio.NewReader(os.Stdin)
    var s int
    fmt.Fscanln(in, &s)
	h := s / 3600 // 3600 секунд в часе
	m := (s - (h * 3600)) / 60 // 60 секунд в минуте
    fmt.Printf("It is %d hours %d minutes.", h, m) 
}

func better() {
    var s uint32
    fmt.Scanln(&s)
    fmt.Printf("It is %d hours %d minutes.", s / 3600, s % 3600 / 60)
}