package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	var rec func(i int) string
	rec = func(i int) string {
		if i >= len(vet) {
			return ""
		}
		if i == len(vet) -1 {
			return fmt.Sprintf("%d", vet[i])
		}
		return fmt.Sprintf("%d, ", vet[i]) + rec(i+1)
	}
	return "[" + rec(0) + "]"
}

func tostrrev(vet []int) string {
	var rec func(i int) string
	rec = func(i int) string {
		if i < 0 {
			return ""
		}
		if i == 0 {
			return fmt.Sprintf("%d", vet[i])
		}
		return fmt.Sprintf("%d, ", vet[i]) + rec(i-1)
	}
	return "[" + rec(len(vet)-1) + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	var rec func(i int)
	rec = func(i int) {
		if i >= len(vet)/2 {
			return
		}
		vet[i], vet[len(vet)-1-i] = vet[len(vet)-1-i], vet[i]
		rec(i + 1)
	}
	rec(0)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	var red func(i int) int
	red = func(i int) int {
		if i >= len(vet) {
			return 0
		}
		return vet[i] + red(i+1)
	}
	return red(0)
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	var red func(i int) int
	red = func(i int) int {
		if i >= len(vet) {
			return 1
		}
		return vet[i] * red(i+1)
	}
	return red(0)
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}
	var rec func(i int) (int, int)

	rec = func(i int) (int, int) {
		if i == len(vet)-1 {
			return vet[i], i
		}
		minVal, minIdx := rec(i + 1)
		if vet[i] < minVal {
			return vet[i], i
		}
		return minVal, minIdx
	}

	_, indice := rec(0)
	return indice
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
