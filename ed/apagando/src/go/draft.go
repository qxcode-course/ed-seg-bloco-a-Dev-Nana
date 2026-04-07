package main
import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m)

	numeros := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&numeros[i])
	}

	fmt.Scan(&n)

	apagar := make(map[int]bool)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		apagar[x] = true
	}

	var resultado []int

	for _, v := range numeros {
		if !apagar[v] {
			resultado = append(resultado, v)
		}
	}

	for _, v := range resultado {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
