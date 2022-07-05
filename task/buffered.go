package main
import "fmt"
 
// Можно меньше но больше нельзя положить в канал или взять 
func main() {
     
    intCh := make(chan int, 3) 
    intCh <- 10
    intCh <- 3
    intCh <- 24
    //intCh <- 15  // блокировка - функция main ждет, когда освободится место в канале
     
    fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	intCh <- 5
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
    fmt.Println("The End")
}

 
func main1() {
     
    intCh := make(chan int, 3) 
    intCh <- 10
     
    fmt.Println(cap(intCh))     // 3
    fmt.Println(len(intCh))     // 1
     
    fmt.Println(<-intCh)
}