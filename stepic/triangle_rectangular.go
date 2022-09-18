package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника.
Нужно проверить, является ли треугольник прямоугольным.
Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
*/
func main1() {
	in := bufio.NewReader(os.Stdin)
	var a, b, c int
	fmt.Fscanln(in, &a, &b, &c)
	if a*a+b*b == c*c {
		fmt.Printf("Прямоугольный")
	} else {
		fmt.Printf("Непрямоугольный")
	}
}

/*
Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	var a, b, c int
	fmt.Fscanln(in, &a, &b, &c)
	approval1 := a + b > c
	approval2 := a + c > b
	approval3 := b + c > a

	if approval1 && approval2 && approval3 {
		fmt.Printf("Существует")
	} else {
		fmt.Printf("Не существует")
	}
}