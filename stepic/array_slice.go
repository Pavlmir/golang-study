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
