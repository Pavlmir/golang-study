package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func task(ctx context.Context) {
	// запускаем бесконечный цикл
	for {
		select {
		// проверяем не завершён ли ещё контекст и выходим, если завершён
		case <-ctx.Done():
			return

		// выполняем нужный нам код
		default:
			println("Hello gophers!")
		}
		// делаем паузу перед следующей итерацией
		time.Sleep(time.Second)
	}
}

func main() {
	// создаём контекст с функцией завершения
	ctx, cancel := context.WithCancel(context.Background())
	// запускаем нашу горутину
	go task(ctx)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Test")
	})

	http.ListenAndServe(":8181", nil)

	// завершаем контекст, чтобы завершить горутину
	cancel()
}
