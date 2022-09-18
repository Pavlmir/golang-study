package main
import "fmt"
 
// функция main ожидает завершения выполнения горутины
func main() {
     
    intCh := make(chan int)  // это небуферизованный канал, ожидаем считывание из канала
     
    go func(){
            fmt.Println("Go routine starts")
            intCh <- 5 // блокировка, пока данные не будут получены функцией main
			intCh <- 6 // запись в канал
			intCh <- 7
			intCh <- 8
    }()
    fmt.Println(<-intCh) // получение данных из канала - main продолжает выполняться и блокируется на этой строке fmt.Println(<-intCh), пока не будут получены данные.
	//intCh <- 9 // блокировка
	fmt.Println(<-intCh) // считываем данные
	fmt.Println(<-intCh) 
    fmt.Println("The End")
}