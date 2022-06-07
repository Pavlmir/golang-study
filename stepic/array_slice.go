package main

import "fmt"

func main() {
    var workArray [10]uint8
	var in, in2 uint8

	for idx := range workArray {
		fmt.Scan(&workArray[idx])
	}

	for idx := 0; idx < 3; idx++ {
		fmt.Scan(&in, &in2)
		workArray[in], workArray[in2] = workArray[in2], workArray[in]
	}
	for _, el := range workArray {
		fmt.Print(el, " ")
	}
}

func task1_12_13(){
	var n int
    fmt.Scan(&n)
    a := make([]int, n, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&a[i])
    }
    fmt.Println(a[3])
}

func max()  {
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
    max := array[0]
    for _, el := range array {
        if el > max {
            max = el
        }
    }
    fmt.Println(max)
}

func task1_12_15() {
    var n int
    fmt.Scan(&n)
    a := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&a[i]) 
        if i % 2 == 0 {
            fmt.Print(a[i], " ")
        }
    }
}

