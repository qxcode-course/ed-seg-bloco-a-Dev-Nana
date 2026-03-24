package main

import "fmt"

func main() {
	var qtd, baruel int
	fmt.Scan(&qtd, &baruel)

	figurinhas := make([]int, baruel)

	for i := 0; i < baruel; i++ {
		fmt.Scan(&figurinhas[i])
	}

	var repetidas []int
	var faltantes []int

	for i := 1; i <= qtd; i++ {
		count := contarNumero(figurinhas, i)

		if count == 0 {
			faltantes = append(faltantes, i)
		}

		if count > 1 {
			for j := 0; j < count-1; j++ {
				repetidas = append(repetidas, i)
			}
		}
	}

	if len(repetidas) > 0 {
		printFormatado(repetidas)
	} else {
		fmt.Println("N")
	}

	if len(faltantes) > 0 {
		printFormatado(faltantes)
	} else {
		fmt.Println("N")
	}
}

func contarNumero(array []int, numero int) int {
	count := 0
	for i := 0; i < len(array); i++ {
		if array[i] == numero {
			count++
		}
	}
	return count
}

func printFormatado(lista []int) {
	for i := 0; i < len(lista); i++ {
		if i != len(lista)-1 {
			fmt.Printf("%d ", lista[i])
		} else {
			fmt.Printf("%d\n", lista[i])
		}
	}
}