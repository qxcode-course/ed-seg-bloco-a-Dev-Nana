package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	nl := len(grid)
	nc := len(grid[0])

	var dfs func(l, c, index int) bool

	dfs = func(l, c, index int) bool {

		if index == len(word) {
			return true
		}

		if l < 0 || l >= nl || c < 0 || c >= nc {
			return false
		}

		if grid[l][c] != word[index] {
			return false
		}

		temp := grid[l][c]
		grid[l][c] = '#'

		if dfs(l-1, c, index+1) || dfs(l+1, c, index+1) || dfs(l, c-1, index+1) || dfs(l, c+1, index+1) {
			grid[l][c] = temp
			return true
		}

		grid[l][c] = temp

		return false
	}

	for i := 0; i < nl; i++{
		for j := 0; j < nc; j++{
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	_, _ = grid, word
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

