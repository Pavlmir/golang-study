package main

import (
	"fmt"
)

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)

}

func sumInt(a ...int) (int, int) {
	sum := 0
	for _, el := range a {
		sum += el
	}
	return len(a), sum
}
