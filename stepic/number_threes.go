package main

import (
	"fmt"
    "strconv"
)

func main() {
	main_ok()
}

func main_my() {
    var numb string
    fmt.Scan(&numb)
    sum := 0
    for i := 0; i<len(numb); i++ {
        a, _ := strconv.Atoi(numb[i:i+1])  
        sum += a     
    }
    fmt.Println(sum)
}

func main_ok() {
    var a int
    fmt.Scan(&a)
    sum := 0
    for a > 0 {
		fmt.Println(a % 10)
        sum += a % 10
		
		
        a /= 10
		fmt.Println(a)
    }
    fmt.Println(sum)
}