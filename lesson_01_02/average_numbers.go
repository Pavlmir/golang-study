package main

import "fmt"

func main() {
	fmt.Println(average_numbers(1, 2, 3, 4))
}

func average_numbers(a float32, b float32, c float32, d float32) float32 {
	return (a + b + c + d) / 4
}
