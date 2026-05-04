package main

import "fmt"

func processa(n int) [][2]int {
	if n == 0 {
		return nil
	}

	result := processa(n / 2)
	return append(result, [2]int{n / 2, n % 2})
}

func main() {
	var n int
	fmt.Scan(&n)

	result := processa(n)

	for _, pair := range result {
		fmt.Println(pair[0], pair[1])
	}
}