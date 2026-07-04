package main
import (
    "fmt"
    "strings"
)

func diagonal(s string, k int) {
    if len(s) == 0 {
        return
    }
    fmt.Println(strings.Repeat(" ", k) + string(s[0]))
    diagonal(s[1:], k+1)
}

func main() {
    var s string
    fmt.Scan(&s)
    diagonal(s, 0)
}