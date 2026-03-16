package main
import "fmt"

type Jogada struct{
    pa, pb int
}

func main() {
    // lista := make([]Jogada, qtd)
    qtd := 0
    fmt.Scan(&qtd)
    jogadas := make([]Jogada, qtd)
    for _, jog := range jogadas{
        // var a, b 
        // fmt.Scan(&a, &b)
        // jogadas = append(jogadas, Jogada{a, b})
        fmt.Scan(&jog.pa, &jog.pb)
    }
    fmt.Println(jogadas)
}
