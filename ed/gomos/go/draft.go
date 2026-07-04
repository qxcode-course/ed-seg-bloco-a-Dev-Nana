package main
import "fmt"

type Pos struct {
    x, y int
}

func main() {
    
    var q int
    var dir string

    fmt.Scan(&q, &dir)

    snake := make([]Pos, q)

    for i := 0; i < q; i++ {
        fmt.Scan(&snake[i].x, &snake[i].y)
    }

    for i := q - 1; i > 0; i--{
        snake[i] = snake[i - 1]
    }

    var x, y int
    fmt.Scan(&x, &y)

    switch dir{
        case "L":
            snake[0].x--
        case "R":
            snake[0].x++
        case "U":
            snake[0].y--
        case "D":
            snake[0].y++
    }

    for i := 0; i < q; i++{
        fmt.Println(snake[i].x, snake[i].y)
    }
}
