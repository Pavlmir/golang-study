package main

import (
	"fmt"
)

func main() {
	//task1_10()
	fmt.Print(1)
	var x, y string
	fmt.Scan(&x, &y)
	fmt.Print(x, y)
}

func task1_10() {
	var x, y string
	fmt.Scan(&x, &y)

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			if x[i] == y[j] {            
				fmt.Print(y[j:j+1], " ")
			}
		}
	}
}
