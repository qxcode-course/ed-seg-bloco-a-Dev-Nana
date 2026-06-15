package main
import "fmt"

func primo (n int, div int) bool{
    if n < 2{
        return false
    }
    if div == 1{
        return true
    }
    if n%div == 0{
        return false
    }
    return primo (n, div-1)
}

func main() {
    var n int
    fmt.Scan(&n)
    primos := []int{}
    atual := 2

    for len(primos) < n{
        if primo(atual, atual-1) {
            primos = append(primos, atual)
        }
        
        atual ++
    }

    fmt.Print("[")
    for i, p := range primos {
        if i > 0 {
            fmt.Print(", ")
        }
        fmt.Print(p)
    }
    fmt.Println("]")
}