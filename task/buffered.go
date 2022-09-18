package main
import "fmt"
 
// Можно меньше но больше нельзя положить в канал или взять 
func main() {
     
    intCh := make(chan int, 3) // это буферизованный канал с емкостью 3, не ожидает извлечения из каналов 
    go func(){
        fmt.Println("Go routine starts")
        intCh <- 10 // передача данных в канал
        intCh <- 3
        intCh <- 24
        //intCh <- 15  // блокировка - функция main ждет, когда освободится место в канале
    }()

    fmt.Println(<-intCh) // получение данных из канала = 10
	fmt.Println(<-intCh) // 3
	intCh <- 5
	fmt.Println(<-intCh) // 24
	fmt.Println(<-intCh) // 5
    fmt.Println("The End")
}

 
func main1() {
     
    intCh := make(chan int, 3) 
    intCh <- 10
     
    fmt.Println(cap(intCh))     // 3
    fmt.Println(len(intCh))     // 1
     
    fmt.Println(<-intCh)
}