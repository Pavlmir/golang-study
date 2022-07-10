/*
Уникальное число
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

func main() {
	in := bufio.NewReader(os.Stdin) // Консоль
	var numbers []int
	for {
		var num_in int
        fmt.Fscanln(in, &num_in)
		if num_in == 0 {
			break
		}
		numbers = append(numbers, num_in)
	}
	middle := len(numbers)/2
	result := 0
	for idx, num := range numbers {
		if idx < middle {
			result += num
		} else {
			result -= num
		}
	}
	fmt.Println(math.Abs(float64(result)))
}