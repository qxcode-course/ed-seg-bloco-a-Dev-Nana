package main
import "fmt"

func subset(vet []int, pos int, target int) bool {
    if target == 0 {
        return true
    }

    if pos >= len(vet){
        return false
    }

    return subset(vet, pos+1, target-vet[pos]) || subset(vet, pos+1, target)
}

func main() {
    var n, m int

    fmt.Scan(&n, &m)

    vet := make([]int, n)

    for i := 0; i < n; i++{
        fmt.Scan(&vet[i])
    }

    fmt.Println(subset(vet, 0, m))
}
