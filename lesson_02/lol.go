package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(generateText(15, 10, 7, 3))

}

func generateText(a int, b int, c int, d int) string {

	var text string = "Всем привет. Мне " + strconv.Itoa(a) + " лет. Я учусь в " + strconv.Itoa(b) + " классе. Я занимаюсь спортом " + strconv.Itoa(c) + " лет. У меня есть брат, ему " + strconv.Itoa(d) + " года"

	return text
}
