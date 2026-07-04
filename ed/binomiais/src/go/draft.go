package main

import "fmt"

func coef(n, k int) int {
    if k == 0 || k == n {
        return 1
    }
    return coef(n-1, k-1) + coef(n-1, k)
}

func main() {
    var n, k int

    fmt.Scan(&n, &k)

    fmt.Println(coef(n, k))
}