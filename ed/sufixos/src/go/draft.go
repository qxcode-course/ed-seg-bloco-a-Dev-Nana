package main
import (
    "fmt"
)

func sufixo(s string){
    if len(s) == 0 {
        return
    }

    sufixo(s[1:])
    fmt.Println(s)
}

func main() {
    var s string
    fmt.Scan(&s)

    sufixo(s)
}