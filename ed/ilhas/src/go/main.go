package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	
	nl := len(grid)
	nc := len(grid[0])

	var dfs func(int, int)

	dfs = func(l, c int) {
		if l < 0 || l >= nl || c < 0 || c >= nc {
			return
		}

		if grid[l][c] != '1' {
			return
		}

		grid[l][c] = '0'

		dfs(l - 1, c)
		dfs(l + 1, c)
		dfs(l, c - 1)
		dfs(l, c + 1)
	}

	ilhas := 0

	for i := 0; i < nl; i++{
		for j := 0; j < nc; j++{
			if grid[i][j] == '1' {
				ilhas++
				dfs(i, j)
			}
		}
	}
	
	return ilhas
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
