package main

import (
    "fmt"
    "math"
)

func main() {
	var rodadas int
	fmt.Scanln(&rodadas)
	vencedor := math.MaxFloat64
    num := 0
	for i := 0; i < rodadas; i++ {
		a := 0
        b := 0
        fmt.Scan(&a, &b)
        if a < 10 || b < 10 {
            continue
        }

        if math.Abs(float64(a - b)) < vencedor {
            num = i
        }
	}
    if num == 0 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(num)
    }
}
