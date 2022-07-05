package main

import (
	"fmt"
)

func main() {
    var a float64
    fmt.Scan(&a)
    if a <= 0 {
        fmt.Printf("число %3.2f не подходит", a) // ширина - 2 и точность - 2, exmpl -000012.2123
    } else if a >= 10_000 {
        fmt.Printf("%e", a) // числа с плавающей точкой в экспоненциальном представлении
    } else {
        fmt.Printf("%.4f", a*a) // ширина по умолчанию и точность - 2 символа
    }
}

/*
функция Sprintf() которая работает как и Printf(), за исключением того что она ничего не печатает,
 а возвращает результат форматирования, рассмотрим пример:
 */

 func exmplSprintf() {
	var a float64 = 100.123456789
	result := fmt.Sprintf("%.2f", a)
	fmt.Printf("%q", result) // вывод: "100.12"
    // result будет типа string
}