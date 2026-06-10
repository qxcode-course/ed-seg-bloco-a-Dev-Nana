package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	count := map[int]int{}

	for _, x := range vet{
		stress := x
		if stress < 0 {
			stress = -stress
		}

		count[stress]++
	}

	keys := []int{}
	for k := range count {
		keys = append(keys, k)
	}
	
	sort.Ints(keys)
	result := []Pair{}

	for _, k := range keys {
		result = append(result, Pair{k, count[k]})
	}

	return result
}

func teams(vet []int) []Pair {
	if len(vet) == 0{
		return []Pair{}
	}

	result := []Pair{}
	current := vet[0]
	count := 1

	for i := 1; i < len(vet); i++{
		if vet[i] == current {
			count++
		} else {
			result = append(result, Pair{current, count})
			current = vet[i]
			count = 1
		}
	}

	result = append(result, Pair{current, count})
	return result
}

func mnext(vet []int) []int {
	result := make([]int, len(vet))

	for i := 0; i < len(vet); i++{
		if vet[i] <= 0{
			continue
		}

		left := i > 0 && vet[i-1] < 0
		right := i < len(vet)-1 && vet[i+1] < 0

		if left || right{
			result[i] = 1
		}
	}
	
	return result
}

func alone(vet []int) []int {
	result := make([]int, len(vet))

	for i := 0; i < len(vet); i++{
		if vet[i] <= 0 {
			continue
		}

		left := i > 0 && vet[i - 1] < 0
		right := i < len(vet) - 1 && vet[i+1] < 0

		if !left && !right {
			result[i] = 1
		}
	}

	return result
}

func couple(vet []int) int {
	men := map[int]int{}
	women := map[int]int{}

	for _, x := range vet {
		if x > 0 {
			men[x]++
		} else {
			women[-x]++
		}
	}

	total := 0

	for stress, qtdMen := range men {
		qtdWomen := women[stress]

		if qtdMen < qtdWomen {
			total += qtdMen
		} else {
			total += qtdWomen
		}
	}
	
	return total
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos+len(seq) > len(vet) {
		return false
	}

	for i := 0; i < len(seq); i++{
		if vet[pos+i] != seq[i] {
			return false
		}
	}

	return true
}

func subseq(vet []int, seq []int) int {
	for pos := 0; pos < len(vet); pos++ {
		if hasSubseq(vet, seq, pos) {
			return pos
		}
	}

	return -1
}

func erase(vet []int, posList []int) []int {
	remove := map[int]bool{}

	for _, pos := range posList {
		remove[pos] = true
	}

	result := []int{}

	for i, value := range vet{
		if !remove[i] {
			result = append(result, value)
		}
	}
	
	return result
}

func clear(vet []int, value int) []int {
	result := []int{}

	for _, x := range vet {
		if x != value {
			result = append(result, x)
		}
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
