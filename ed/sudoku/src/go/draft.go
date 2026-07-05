package main

import (
    "fmt"
    "bufio"
    "os"
)

func existeLinha(matriz [][]rune, lin int, valor rune) bool {
    for _, x := range matriz[lin] {
        if x == valor {
            return true
        }
    }
    return false
}

func existeColuna(matriz [][]rune, col int, valor rune) bool {
    for i := 0; i < len(matriz); i++ {
        if matriz[i][col] == valor {
            return true
        }
    }
    return false
}

func quadrante(matriz [][]rune, lin, col int) []rune {
    tam := 3
    if len(matriz) == 4{
        tam = 2
    }

    l := (lin / tam) * tam
    c := (col / tam) * tam

    valores := []rune{}

    for i := 0; i < tam; i++{
        for j := 0; j < tam; j++{
            valores = append(valores, matriz[l + i][c+j])
        }
    }

    return valores
}

func existeQuadrante(matriz [][]rune, lin, col int, valor rune) bool{
    for _, x := range quadrante(matriz, lin, col) {
        if x == valor {
            return true
        }
    }
    return false
}

func resolver(matriz [][]rune, index int) bool {
    n := len(matriz)

    if index == n*n {
        return true
    }

    lin := index / n
	col := index % n

    if matriz[lin][col] != '.' {
        return resolver(matriz, index+1)
    }

    for num := '1'; num <= rune('0'+n); num++{
        if !existeLinha(matriz, lin, num) && !existeColuna(matriz, col, num) && !existeQuadrante(matriz, lin, col, num) {
            matriz[lin][col] = num

            if resolver(matriz, index+1){
                return true
            }

            matriz[lin][col] = '.'
        }
    }
    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var n int

    scanner.Scan()
    
    fmt.Sscan(scanner.Text(), &n)
    matriz := make([][]rune, n)

    for i := 0; i < n; i++{
        scanner.Scan()
        matriz[i] = []rune(scanner.Text())
    }

    resolver(matriz, 0)
    for i := 0; i < n; i++{
        fmt.Println(string(matriz[i]))
    }
}