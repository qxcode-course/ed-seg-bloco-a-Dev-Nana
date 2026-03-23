package main

import "fmt"

func main() {
	var qtd int
	count := 0

	fmt.Scanln(&qtd)
	animais := make([]int, qtd)

	for i := 0; i < qtd; i++ {
		fmt.Scan(&animais[i])
	}

	for i := 0; i < len(animais)-1; i++ {
		if (animais[i] < 0 && animais[i+1] > 0) ||
			(animais[i] > 0 && animais[i+1] < 0) {

			animais[i] = 0
			animais[i+1] = 0
			count++
		}
	}

	fmt.Println(count)
}