package main

import "fmt"

func escada(n int) int {
    if n == 1{
        return 1
    }

    if n == 2 {
        return 1
    }

    if n == 3{
        return 2
    }

    return escada(n-1) + escada(n-3)
}

func main() {
    var n int 
    fmt.Scan(&n)

    fmt.Println(escada(n))
}