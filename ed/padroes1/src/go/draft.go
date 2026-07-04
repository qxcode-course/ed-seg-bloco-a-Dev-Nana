package main
import "fmt"

func bloco(n int) int{
    if n == 1{
        return 20
    }
    return bloco(n-1) + 8
}

func main(){
    var n int
    fmt.Scan(&n)
    fmt.Println(bloco(n))
}