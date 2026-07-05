package main

import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    expr := scanner.Text()

    pilha := []rune{}

    for _, c := range expr {

        if c == '(' || c == '[' {
            pilha = append(pilha, c)
        } else {
            if len(pilha) == 0 {
                fmt.Println("nao balanceado")
                return
            }
            
            topo := pilha[len(pilha)-1]
            pilha = pilha[:len(pilha)-1]

            if (c == ')' && topo != '(') || (c == ']' && topo != '[') {
                fmt.Println("nao balanceado")
                return
            }
        }
    }

    if len(pilha) == 0 {

        fmt.Println("balanceado")
        
    } else {
        
        fmt.Println("nao balanceado")

    }
}