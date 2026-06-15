package main
import "fmt"
func main() {
    
    var q int
    var dir string

    fmt.Scan(&q, &dir)

    var x, y int
    fmt.Scan(&x, &y)

    switch dir{
        case "L":
            x--
        case "R":
            x++
        case "U":
            y--
        case "D":
            y++
    }

    fmt.Println(x, y)
}
