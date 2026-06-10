package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(arr []int, value int) (bool, int) {
	low := 0
	high := len(arr) 

	for low < high {
		mid := (low + high) / 2

		if arr[mid] == value {
			return true, mid
		}
		
		if value < arr[mid] {
			high = mid 
		} else {
			low = mid + 1
		}
	}
	return false, low
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
