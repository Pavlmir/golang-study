package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	intCh := make(chan int)

	go factorial(7, intCh, ctx)

	for num := range intCh {
		fmt.Println(num)
	}
}

func factorial(n int, ch chan int, ctx context.Context) {
	defer close(ch)
	// result := 1
	// for i := 1; i <= n; i++ {
	// 	result *= i
	// 	ch <- result
	// }
	// запускаем бесконечный цикл
	for {
		select {
		// проверяем не завершён ли ещё контекст и выходим, если завершён
		case <-ctx.Done():
			return

		// выполняем нужный нам код
		default:
			println("Hello gophers!", n)
		}
		// делаем паузу перед следующей итерацией
		time.Sleep(time.Second)
	}
}
