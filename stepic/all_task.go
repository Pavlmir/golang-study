package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var a, b float64
	fmt.Fscanln(in, &a, &b)
	fmt.Println((a + b) / 2)
}
