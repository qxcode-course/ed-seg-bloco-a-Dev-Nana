package main
import "fmt"

func solve(colunas []int, linha int, n int) int{
    if linha == n{
        return 1
    }
    total := 0
    
    for coluna := 0; coluna < n; coluna++{
        v := true

        for i := 0; i < linha; i++{
        
        if colunas[i] == coluna{
            v = false
        }

        diffCol := colunas[i] - coluna
        if diffCol < 0 {
            diffCol = -diffCol
        }

        diffLin := i - linha
        
        if diffLin < 0 {
            diffLin = -diffLin
        }

        if diffCol == diffLin {
            v = false
        }
    }
        if v {
            colunas[linha] = coluna
            total += solve(colunas, linha+1, n)
        }
    }


    return total
}

func main() {
    var n int
    fmt.Scan(&n)

    colunas := make([]int, n)

    fmt.Println(solve(colunas, 0, n))
}