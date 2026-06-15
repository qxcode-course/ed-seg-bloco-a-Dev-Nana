package main
import "fmt"

func primo(n int, div int) bool{
    if n < 2{
        return false
    }
    if div == 1{
        return true
    }
    if n%div == 0{
        return false
    }
    return primo(n, div-1)
}

func esimo(n, atual, cont int) int {
    if primo(atual, atual-1){
        cont++
    }

    if cont == n{
        return atual
    }

    return esimo(n, atual+1, cont)
}

func main() {
    var n int 
    fmt.Scan(&n)
    
    fmt.Println(esimo(n, 2, 0))
}