/*
Замыкание привязывает к себе внешнюю переменную.
После выхода из внешней функции `Generate` она не уничтожается, а остаётся привязанной к функции замыкания,
причём её значение сохраняется между вызовами функции.

Такие функции иногда называют генераторами.
*/
package main

import (
	"fmt"
)

func Generate(seed int) func() {
	return func() {
		fmt.Println(seed) // замыкание получает внешнюю переменную seed
		seed += 2         // переменная модифицируется
	}

}

func main() {
	iterator := Generate(0)
	iterator()
	iterator()
	iterator()
	iterator()
	iterator()
}
