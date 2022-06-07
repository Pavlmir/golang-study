package main

import "fmt"

func main(){
    var n1, n2 int
    fmt.Scan(&n1, &n2)
    fmt.Println(max(n1, n2))
    
}

func max(n1 int, n2 int) int {
    if n1 > n2 {
        return n1
    }
    return n2
}

func calc(a int) (int, int) {
    return a * 2, a * a
}

func isEven(a int) bool {
    return a % 2 == 0
}

func double_m(a, b int) int {
    p := 0
    for i := a; i <= b; i++ {
        p = p + i * i  
    }
    return p  
}