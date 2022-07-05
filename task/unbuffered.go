package main
import "fmt"
 
// функция main ожидает завершения выполнения горутины
func main() {
     
    intCh := make(chan int) 
     
    go func(){
            fmt.Println("Go routine starts")
            intCh <- 5 // блокировка, пока данные не будут получены функцией main
			intCh <- 6
			intCh <- 7
			intCh <- 8
    }()
    fmt.Println(<-intCh) // получение данных из канала - main продолжает выполняться и блокируется на этой строке fmt.Println(<-intCh), пока не будут получены данные.
	// intCh <- 9
	fmt.Println(<-intCh) 
	fmt.Println(<-intCh) 
    fmt.Println("The End")
}