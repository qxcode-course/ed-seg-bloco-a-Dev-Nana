package main
import (
    "fmt"
    "bufio"
    "os"
)

func valido(seq []rune, pos int, valor rune, limite int) bool{
    for i := pos - limite; i <= pos+limite; i++ {
        if i < 0 || i >= len(seq) || i == pos {
            continue
        }
        if seq[i] == valor {
            return false
        }
    }
    return true
}

func resolver(seq []rune, limite, index int) bool{
    if index == len(seq){
        return true
    }

    if seq[index] != '.' {
        return resolver(seq, limite, index+1)
    }

    for num := 0; num <= limite; num++ {
        valor := rune('0' + num)

        if valido(seq, index, valor, limite){
            seq[index] = valor

            if resolver(seq, limite, index+1) {
                return true
            }
            seq[index] = '.'
        }
    }

    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    seq := []rune(scanner.Text())

    var limite int
    scanner.Scan()
    fmt.Sscan(scanner.Text(), &limite)

    resolver(seq, limite, 0)

    fmt.Println(string(seq))
}
